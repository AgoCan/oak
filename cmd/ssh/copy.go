package ssh

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/agocan/oak/pkg/data"
	"github.com/agocan/oak/pkg/ssh"
)

var _recursion bool

func CopyCmd() *cobra.Command {

	var command = &cobra.Command{
		Use:   "cp [machine/group]",
		Short: "Excute command.",
		Long:  ``,
		// Args:  cobra.MinimumNArgs(2),
		Run: Copy,
	}
	command.PersistentFlags().BoolVarP(&_recursion, "recursion", "r", false, "Recursion director.")
	return command
}

func Copy(cmd *cobra.Command, args []string) {
	group := data.GetGroup(args[0])
	if group != nil {
		command := strings.Join(args[1:], " ")
		ssh.GroupExec(group, command)
		return
	}

	machine := data.GetMachine(args[0])
	if machine == nil {
		logMsg := fmt.Sprintf("%s machine not found.\n", args[0])
		log.Fatal(logMsg)
	}
	sshClient := ssh.NewSsh(machine)

	sshClient.Copy()
}
