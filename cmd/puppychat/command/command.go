package command

import (
    "github.com/spf13/cobra"
    "log"
)

func command() *cobra.Command {
    cmd := &cobra.Command{
        Use: "chatpuppy",
        CompletionOptions: cobra.CompletionOptions{
            DisableDefaultCmd:   true,
            DisableNoDescFlag:   true,
            DisableDescriptions: true,
            HiddenDefaultCmd:    true,
        },
    }

    cmd.AddCommand(
        apiCommand(),
        serialCommand(),
        nodeCommand(),
    )

    return cmd
}

func Run() {
    err := command().Execute()
    if err != nil {
        log.Fatal(err)
    }
}
