package cmd

import (
	"github.com/spf13/cobra"

	"github.com/agocan/oak/cmd/data"
	"github.com/agocan/oak/cmd/ssh"
)

var RootCmd = &cobra.Command{Use: "oak"}

func init() {

	RootCmd.AddCommand(data.MachineCommand())
	RootCmd.AddCommand(data.GroupCommand())

	RootCmd.AddCommand(ssh.TerminalCommand())

}
