package g

import (
    "fmt"
    "github.com/chatpuppy/puppychat/assets"
    "github.com/fsnotify/fsnotify"
    log "github.com/sirupsen/logrus"
    "github.com/spf13/viper"
    "os"
    "path/filepath"
)

type ConfigApp struct {
    Address      string
    Mode         string
    TZ           string
    PrintStack   bool
    ApiUrl       string
    Distribution ConfigDistribution
}

type ConfigDistribution struct {
    IP   string
    Port int
}

func (d ConfigDistribution) String() string {
    return fmt.Sprintf("%s:%d", d.IP, d.Port)
}

type ConfigDB struct {
    Dsn   string
    Debug bool
}

type ConfigRedis struct {
    Address  string
    DB       int
    Password string
}

type ConfigLogger struct {
    Color    bool
    Level    string
    Age      int
    Json     bool
    RootPath string
}

type config struct {
    App      ConfigApp
    Database ConfigDB
    Redis    ConfigRedis
    Logger   ConfigLogger
}

var (
    Config     *config
    configFile string
)

func SetConfigFile(path string) {
    configFile = path
}

func readConfig() error {
    viper.SetConfigFile(configFile)
    viper.AutomaticEnv()
    err := viper.ReadInConfig()
    if err != nil {
        return err
    }
    Config = new(config)
    err = viper.Unmarshal(Config)
    return err
}

// InitializeConfig initialize configure
func InitializeConfig() {
    // Check config file if exists
    if _, err := os.Stat(configFile); os.IsNotExist(err) {
        // create directory
        d := filepath.Dir(configFile)
        if _, err = os.Stat(d); os.IsNotExist(err) {
            err = os.MkdirAll(d, 0755)
            if err != nil {
                log.Fatalf("Config directory create failed: %v", err)
            }
        }
        // create configure file use default
        err = os.WriteFile(configFile, assets.DefaultConfig, 0644)
        if err != nil {
            log.Fatalf("Config file save failed: %v", err)
        }
    }

    err := readConfig()
    if err != nil {
        log.Fatalf("Config read failed: %v", err)
    }

    // setting apiurl
    node.ApiUrl = Config.App.ApiUrl

    // Config file modify monitor
    viper.OnConfigChange(func(e fsnotify.Event) {
        // log.Infof("%s changed, reloading", e.Name)
    })

    viper.WatchConfig()
}
