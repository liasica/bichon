package command

import (
    "context"
    "fmt"
    "github.com/chatpuppy/puppychat/app/router"
    "github.com/chatpuppy/puppychat/grpc/pb/message"
    pb "github.com/chatpuppy/puppychat/grpc/pb/node"
    "github.com/chatpuppy/puppychat/internal"
    "github.com/chatpuppy/puppychat/internal/g"
    "github.com/chatpuppy/puppychat/utils"
    "github.com/spf13/cobra"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    "google.golang.org/grpc/keepalive"
    "google.golang.org/grpc/metadata"
    "io"
    "log"
    "time"
)

var (
    NodeSecret string
)

func apiCommand() *cobra.Command {

    var (
        addr       string
        configFile string
        rootPath   string
    )

    cmd := &cobra.Command{
        Use:   "server",
        Short: "Run api server",
        Run: func(_ *cobra.Command, _ []string) {
            err := g.IinitializeNode(NodeSecret)
            if err != nil {
                log.Fatalf("Node initialize failed: %v", err)
            }

            // server settings
            g.SetConfigFile(configFile)
            internal.SetRootPath(rootPath)
            internal.Bootstrap()

            // start api server
            go router.Run()

            // start grpc client
            var kacp = keepalive.ClientParameters{
                Time:                10 * time.Second,
                Timeout:             time.Second,
                PermitWithoutStream: true,
            }

            conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithKeepaliveParams(kacp))
            if err != nil {
                log.Fatalf("can not connect grpc server: %v", err)
            }

            defer func(conn *grpc.ClientConn) {
                _ = conn.Close()
            }(conn)

            c := pb.NewNodeClient(conn)

            // ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
            // defer cancel()
            //
            // var res *message.Response
            // res, err = c.GetNodes(ctx, &message.NodeRequest{Sn: utils.GetMachineID()})
            // if err != nil {
            //     log.Fatalf("unexpected error from GetNodes: %v", err)
            // }
            // fmt.Printf("RPC response: %#v", res.Data)
            // select {}

            SyncMessage(c)
        },
    }

    cmd.PersistentFlags().StringVarP(&addr, "grpc-server", "s", "localhost:5301", "grpc server")
    cmd.PersistentFlags().StringVarP(&configFile, "config", "c", "./config/chatpuppy.yaml", "Config file path")
    cmd.PersistentFlags().StringVar(&rootPath, "rootpath", "", "Source code root path (split logging path)")

    return cmd
}

func SyncMessage(c pb.NodeClient) {
    // Create metadata and context.
    md := metadata.Pairs("timestamp", utils.GetTimestampNanoString())
    ctx := metadata.NewOutgoingContext(context.Background(), md)

    // Make RPC using the context with the metadata.
    stream, err := c.Sync(ctx)
    if err != nil {
        log.Fatalf("failed to call Sync: %v\n", err)
    }

    go func() {
        defer func(stream pb.Node_SyncClient) {
            _ = stream.CloseSend()
        }(stream)

        // Read the header when the header arrives.
        var header metadata.MD
        header, err = stream.Header()
        if err != nil {
            log.Fatalf("failed to get header from stream: %v", err)
        }
        // Read metadata from server's header.
        if t, ok := header["timestamp"]; ok {
            fmt.Printf("timestamp from header:\n")
            for i, e := range t {
                fmt.Printf(" %d. %s\n", i, e)
            }
        } else {
            log.Fatal("timestamp expected but doesn't exist in header")
        }
        // if l, ok := header["location"]; ok {
        //     fmt.Printf("location from header:\n")
        //     for i, e := range l {
        //         fmt.Printf(" %d. %s\n", i, e)
        //     }
        // } else {
        //     log.Fatal("location expected but doesn't exist in header")
        // }

        // Send all requests to the server.
        for {
            if err = stream.Send(&message.SyncRequest{Message: "Test message", Sn: utils.GetMachineID()}); err != nil {
                log.Fatalf("failed to send streaming: %v\n", err)
            }
            time.Sleep(5 * time.Second)
        }
    }()

    // Read all the responses.
    var rpcStatus error
    fmt.Printf("response:\n")
    for {
        var r *message.SyncResponse
        r, err = stream.Recv()
        if err != nil {
            rpcStatus = err
            break
        }
        fmt.Printf(" - %s\n", r.Message)
    }
    if rpcStatus != io.EOF {
        log.Fatalf("failed to finish server streaming: %v", rpcStatus)
    }

    // Read the trailer after the RPC is finished.
    trailer := stream.Trailer()
    // Read metadata from server's trailer.
    if t, ok := trailer["timestamp"]; ok {
        fmt.Printf("timestamp from trailer:\n")
        for i, e := range t {
            fmt.Printf(" %d. %s\n", i, e)
        }
    } else {
        log.Fatal("timestamp expected but doesn't exist in trailer")
    }
}
