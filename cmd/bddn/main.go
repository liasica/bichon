package main

import (
    "github.com/chatpuppy/puppychat/cmd/bddn/service"
    pb "github.com/chatpuppy/puppychat/grpc/pb/node"
    "google.golang.org/grpc"
    "google.golang.org/grpc/keepalive"
    "log"
    "net"
    "time"
)

var kaep = keepalive.EnforcementPolicy{
    MinTime:             5 * time.Second,
    PermitWithoutStream: true,
}

var kasp = keepalive.ServerParameters{
    MaxConnectionIdle:     15 * time.Second,
    MaxConnectionAge:      30 * time.Second,
    MaxConnectionAgeGrace: 5 * time.Second,
    Time:                  5 * time.Second,
    Timeout:               1 * time.Second,
}

func main() {
    lis, err := net.Listen("tcp", ":5301")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp))
    pb.RegisterNodeServer(s, service.NewNode())

    if err = s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
