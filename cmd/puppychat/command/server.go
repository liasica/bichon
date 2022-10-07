package command

import (
    "github.com/chatpuppy/puppychat/app/distribution"
    "github.com/chatpuppy/puppychat/app/model"
    "github.com/chatpuppy/puppychat/app/router"
    "github.com/chatpuppy/puppychat/app/ws"
    "github.com/chatpuppy/puppychat/internal"
    "github.com/chatpuppy/puppychat/internal/g"
    "github.com/chatpuppy/puppychat/pkg/logger"
    log "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
)

var (
    NodeSecret string
)

func serverCommand() *cobra.Command {

    var (
        addr       string
        configFile string
        rootPath   string
    )

    cmd := &cobra.Command{
        Use:   "server",
        Short: "Run api server and distribution server (if needs to)",
        Run: func(_ *cobra.Command, _ []string) {
            logger.LoadWithConfig(logger.Config{
                Color: true,
                Level: "info",
                Age:   8192,
            })

            err := g.IinitializeNode(NodeSecret)
            if err != nil {
                log.Fatalf("Node initialize failed: %v", err)
            }

            // server settings
            g.SetConfigFile(configFile)

            internal.ApiBootstrap()

            // create websocket hub
            ws.CreateHub()

            // if node number is zero
            // create distribution hub
            if g.IsDistributionNode() {
                distribution.CreateHub()
            } else {
                // start sync client
                distribution.CreateClient()
                log.Info("waiting for data synchronization to complete")
                <-model.DistributionNodeSyncedChan
            }

            // start api server
            go router.Run()

            select {}
        },
    }

    cmd.PersistentFlags().StringVarP(&addr, "grpc-server", "s", "localhost:5301", "grpc server")
    cmd.PersistentFlags().StringVarP(&configFile, "config", "c", "./config/chatpuppy.yaml", "Config file path")
    cmd.PersistentFlags().StringVar(&rootPath, "rootpath", "", "Source code root path (split logging path)")

    return cmd
}
