package g

import (
    "fmt"
    log "github.com/sirupsen/logrus"
    "os"
    "os/exec"
    "time"
)

var (
    SkipSyncTime string
    ntpd         string
)

func AutoSyncTime() {
    if SkipSyncTime == "true" {
        return
    }

    var err error
    ntpd, err = exec.LookPath("ntpd")
    if err != nil || ntpd == "" {
        log.Fatalf("ntpd found in $PATH=%s, please install openntpd first", os.Getenv("PATH"))
    }

    log.Infof("found ntpd: %s, starting sync time", ntpd)

    ticker := time.NewTicker(60 * time.Second)
    for ; true; <-ticker.C {
        SyncTime()
    }
}

func SyncTime() {
    _, err := exec.Command("/bin/sh", "-c", fmt.Sprintf("%s -d -n -p pool.ntp.org", ntpd)).Output()
    if err != nil {
        log.Errorf("ntpd update time failed: %v", err)
    }
}
