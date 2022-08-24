package internal

import (
    "fmt"
    "os"
    "strings"
)

const (
    genFile = "package ent\n\n//go:generate go run -mod=mod github.com/chatpuppy/puppychat/cmd/ent generate ./schema\n"
)

// examples formats the given examples to the cli.
func examples(ex ...string) string {
    for i := range ex {
        ex[i] = "  " + ex[i] // indent each row with 2 spaces.
    }
    return strings.Join(ex, "\n")
}

func createDir(target string) error {
    _, err := os.Stat(target)
    if err == nil || !os.IsNotExist(err) {
        return err
    }
    if err = os.MkdirAll(target, os.ModePerm); err != nil {
        return fmt.Errorf("creating schema directory: %w", err)
    }
    if target != defaultSchema {
        return nil
    }
    if err = os.WriteFile("./internal/ent/generate.go", []byte(genFile), 0644); err != nil {
        return fmt.Errorf("creating generate.go file: %w", err)
    }
    return nil
}
