package cmd

import (
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"

	"github.com/agocan/oak/cmd/data"
	"github.com/agocan/oak/cmd/ssh"
	"github.com/agocan/oak/config"
)

var _configFile string

func _rootCmd() *cobra.Command {

	var command = &cobra.Command{
		Use: "oak",
	}
	command.PersistentFlags().StringVar(&_configFile, "c", "", "Config file path.")

	command.AddCommand(data.MachineCommand())
	command.AddCommand(data.GroupCommand())

	command.AddCommand(ssh.TerminalCommand())
	command.AddCommand(ssh.ExecCmd())

	// _rootCmd.AddCommand(ssh.CopyCmd())

	return command
}

func Execute() error {
	cobra.OnInitialize(initConfig)
	return _rootCmd().Execute()
}

func initConfig() {
	if _configFile == "" {
		oakDirector := config.GetOakDefaultDir()
		_configFile = path.Join(oakDirector, "config.yaml")
		if _, err := os.Stat(_configFile); os.IsNotExist(err) {
			f, err := os.OpenFile(_configFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
		}

	}
	config.Init(_configFile)
}
