package config

import (
	"log"
	"path"
	"runtime"

	"github.com/spf13/viper"
)

var (
	ExecLogPath string
	StorageParh string
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
	config := viper.New()
	config.AutomaticEnv()

	config.SetConfigFile(configFile)
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	StorageParh = config.GetString("SRORAGE_PARH")
	if StorageParh == "" {
		oakDirector := GetOakDefaultDir()
		StorageParh = path.Join(oakDirector, "data.yaml")
	}

	ExecLogPath = config.GetString("EXEC_LOG_PATH")
	if ExecLogPath == "" {
		oakDirector := GetOakDefaultDir()
		ExecLogPath = path.Join(oakDirector, "log")
	}
}
