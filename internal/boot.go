package internal

import (
    "github.com/chatpuppy/puppychat/internal/ent"
    "github.com/chatpuppy/puppychat/internal/g"
    "github.com/chatpuppy/puppychat/pkg/logger"
    "github.com/golang-module/carbon/v2"
    "os"
    "time"
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

    // set timezone to UTC
    tz := time.UTC.String()
    _ = os.Setenv("TZ", tz)
    loc, _ := time.LoadLocation(tz)
    time.Local = loc
    carbon.SetTimezone(tz)

    // connect database
    ent.OpenDatabase(g.Config.Database.Dsn, g.Config.Database.Debug)
}
