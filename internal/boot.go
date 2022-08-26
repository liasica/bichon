package internal

import (
    "github.com/chatpuppy/puppychat/internal/ent"
    "github.com/chatpuppy/puppychat/internal/g"
    "github.com/chatpuppy/puppychat/pkg/logger"
)

var (
    rootPath string
)

func SetRootPath(path string) {
    rootPath = path
}

func Bootstrap() {
    // logger
    logger.LoadWithConfig(logger.Config{
        Color:    true,
        Level:    "info",
        Age:      8192,
        RootPath: rootPath,
    })

    // load configure from file
    g.InitializeConfig()

    // connect database
    ent.OpenDatabase(g.Config.Database.Dsn, g.Config.Database.Debug)
}
