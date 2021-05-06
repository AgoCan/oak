package ssh

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/spf13/cobra"

	"github.com/agocan/oak/pkg/data"
	"github.com/agocan/oak/pkg/ssh"
)

var wg sync.WaitGroup

func ExecCmd() *cobra.Command {
	var command = &cobra.Command{
		Use:   "exec [machine/group]",
		Short: "Excute command.",
		Long:  ``,
		Args:  cobra.MinimumNArgs(2),
		Run:   Exec,
	}
	return command
}

func Exec(cmd *cobra.Command, args []string) {
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
	command := strings.Join(args[1:], " ")
	sshClient.Excute(command)
}
