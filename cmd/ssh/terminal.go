package ssh

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/agocan/oak/pkg/data"
	"github.com/agocan/oak/pkg/ssh"
)

func TerminalCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "ssh [command]",
		Short: "Operating groups.",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run:   Terminal,
	}
	return command
}

func Terminal(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		log.Fatal("Please use one machine name.")
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	machine := data.GetMachine(args[0])
	if machine == nil {
		logMsg := fmt.Sprintf("%s machine not found.\n", args[0])
		log.Fatal(logMsg)
	}
	sshClient := ssh.NewSsh(machine)

	go func() {
		if err := sshClient.Terminal(); err != nil {
			log.Print(err)
		}
		sshClient.Cancel()
	}()

	select {
	case <-sig:
		sshClient.Cancel()
	case <-sshClient.Ctx.Done():
	}
}
