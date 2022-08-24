package main

import (
    "github.com/chatpuppy/puppychat/cmd/ent/internal"
    "github.com/spf13/cobra"
)

func main() {
    cmd := &cobra.Command{
        Use: "ent",
    }
    cmd.AddCommand(
        internal.InitCmd(),
        internal.GenerateCmd(),
        internal.CleanCmd(),
    )
    _ = cmd.Execute()
}
