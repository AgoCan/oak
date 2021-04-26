package cmd

import (
	"github.com/agocan/oak/cmd/data"
	"github.com/agocan/oak/cmd/ssh"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{Use: "oak"}

func init() {

	RootCmd.AddCommand(data.MacCmd)
	RootCmd.AddCommand(data.GroupCmd)

	RootCmd.AddCommand(ssh.ExecCmd)
	RootCmd.AddCommand(ssh.CopyCmd)

	RootCmd.AddCommand(ssh.TestConnectCmd)

}
