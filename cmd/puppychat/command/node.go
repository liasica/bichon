package command

import (
    "fmt"
    "github.com/chatpuppy/puppychat/internal/g"
    "github.com/chatpuppy/puppychat/utils"
    "github.com/spf13/cobra"
    "log"
    "math"
    "os"
)

func nodeCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "node",
        Short: "Node tools",
    }
    cmd.AddCommand(nodeGenerate())
    return cmd
}

func nodeGenerate() *cobra.Command {
    var (
        nodeid int64
    )

    cmd := &cobra.Command{
        Use:   "generate",
        Short: "Generate new node's rsa keys",
        Run: func(_ *cobra.Command, _ []string) {
            private, public := utils.NewRsa().GenRsaKey()
            n := &g.Node{
                PublicKey:  public,
                PrivateKey: private,
                NodeID:     nodeid,
            }
            str, err := n.Marshal()
            if err != nil {
                log.Fatalf("Node generate failed: %s", err.Error())
            }
            fmt.Printf("NodeID is:\n%d\n\n", nodeid)
            fmt.Printf("Public key is:\n%s\n\n", public)
            fmt.Printf("Private key is:\n%s\n\n", private)
            fmt.Printf("Node secret is:\n%s\n", str)

            // save file
            d := "config/nodes"
            if _, err = os.Stat(d); os.IsNotExist(err) {
                err = os.MkdirAll(d, 0755)
                if err != nil {
                    log.Fatalf("nodes directory create failed: %v", err)
                }
            }

            _ = os.WriteFile(fmt.Sprintf("%s/%d.node", d, nodeid), []byte(str), 0755)
        },
    }

    cmd.Flags().Int64Var(&nodeid, "nodeid", 0, fmt.Sprintf("Node's appid, max number is: %d", math.MaxUint16))
    _ = cmd.MarkFlagRequired("nodeid")

    return cmd
}
