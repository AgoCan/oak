package config

import (
	"log"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

var (
	ExecLogPath string
	SrorageParh string
)

// Get abs path
func getCurrPath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

// init config
func Init(configFile string) {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		configFile = path.Join(path.Dir(getCurrPath()), "config/config.yaml")
	}

	config := viper.New()
	config.AutomaticEnv()
	config.SetConfigFile(configFile)
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	SrorageParh = config.GetString("SRORAGE_PARH")
	ExecLogPath = config.GetString("EXEC_LOG_PATH")
	if !strings.HasSuffix(ExecLogPath, "/") {
		log.Fatal("EXEC_LOG_PATH does not end with /")
	}
}
