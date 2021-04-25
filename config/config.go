package config

// 配置文件导入yaml文件是configstruct.go
//
// 配置文件可以使用 -c 的参数
// https://github.com/go-yaml/yaml

import (
	"log"
	"os"
	"path"
	"runtime"

	"github.com/spf13/viper"
)

// 设置配置文件的 环境变量
var (
	// LogDirector 日志目录
	SrorageParh string
)

// 获取文件绝对路径
func getCurrPath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

// InitConfig 初始化配置项
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

}
