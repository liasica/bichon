package command

import (
    "fmt"
    "github.com/chatpuppy/puppychat/utils"
    "github.com/spf13/cobra"
)

func serialCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "serial",
        Short: "Print current node serial code",
        Run: func(_ *cobra.Command, _ []string) {
            id := utils.GetMachineID()
            fmt.Printf("Your node serial code is: %s\n", id)
        },
    }
}
