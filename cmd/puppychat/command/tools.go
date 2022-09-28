package command

import (
    "fmt"
    "github.com/golang-module/carbon/v2"
    "github.com/spf13/cobra"
    "time"
)

func toolsCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "tools",
        Short: "Global tools",
    }
    cmd.AddCommand(
        timeTool(),
    )
    return cmd
}

func timeTool() *cobra.Command {
    return &cobra.Command{
        Use:   "time",
        Short: "Show current time",
        Run: func(_ *cobra.Command, _ []string) {
            fmt.Println(time.Now().Format(carbon.ISO8601NanoLayout))
        },
    }
}
