package service

import (
    "context"
    "fmt"
    "github.com/chatpuppy/puppychat/grpc/pb/message"
    pb "github.com/chatpuppy/puppychat/grpc/pb/node"
    "github.com/chatpuppy/puppychat/utils"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/metadata"
    "google.golang.org/grpc/status"
    "io"
    "sync"
    "time"
)

type nodeService struct {
    pb.UnimplementedNodeServer
    nodes sync.WaitGroup
}

func NewNode() *nodeService {
    return &nodeService{}
}

// GetNodes Get all nodes of network
func (*nodeService) GetNodes(context.Context, *message.NodeRequest) (*message.Response, error) {
    return &message.Response{
        Data: &message.Response_Nodes{Nodes: &message.GetNodesResponse{
            Nodes: []*message.NodeItem{
                {
                    Sn:   utils.GetMachineID(),
                    Addr: "addr",
                },
            },
        }},
    }, nil
}

// Register Check if the node is valid and register to network
func (*nodeService) Register(context.Context, *message.NodeRequest) (*message.Response, error) {
    return nil, nil
}

func (*nodeService) Sync(stream pb.Node_SyncServer) error {
    defer func() {
        trailer := metadata.Pairs("timestamp", utils.GetTimestampNanoString())
        stream.SetTrailer(trailer)
    }()

    md, ok := metadata.FromIncomingContext(stream.Context())
    if !ok {
        return status.Errorf(codes.DataLoss, "Sync: failed to get metadata")
    }

    var t []string
    if t, ok = md["timestamp"]; ok {
        fmt.Printf("timestamp from metadata:\n")
        for i, e := range t {
            fmt.Printf(" %d. %s\n", i, e)
        }
    }

    // Create and send header.
    header := metadata.New(map[string]string{"timestamp": time.Now().Format(utils.GetTimestampNanoString())})
    _ = stream.SendHeader(header)

    // Read requests and send responses.
    for {
        in, err := stream.Recv()
        if err == io.EOF {
            return nil
        }
        if err != nil {
            return err
        }
        fmt.Printf("request received %v\n", in)
        if err = stream.Send(&message.SyncResponse{Message: in.Message}); err != nil {
            return err
        }
    }
}
