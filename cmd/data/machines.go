package data

import (
	"log"

	"github.com/spf13/cobra"

	msgData "github.com/agocan/oak/pkg/data"
)

func MachineCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "mac [command]",
		Short: "Operating machine.",
		Long:  ``,
	}
	command.AddCommand(MachineAddCommand())
	command.AddCommand(MachineRemoveCommand())
	command.AddCommand(MachineUpdateCommand())
	command.AddCommand(MachineListCommand())
	return command
}

func MachineAddCommand() *cobra.Command {
	var mac msgData.Machine
	var command = &cobra.Command{
		Use:   "add machine_name",
		Short: "Add machine.",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			mac.Name = args[0]
			if mac.Host == "" {
				log.Fatal("Not host.Use -H add host message.")
			}
			mac.Add()
		},
	}
	command.PersistentFlags().StringVarP(&mac.Host, "host", "H", "", "Host.")
	command.PersistentFlags().StringVarP(&mac.User, "user", "u", "root", "User.")
	command.PersistentFlags().StringVarP(&mac.Password, "password", "p", "123456", "Password.")
	command.PersistentFlags().IntVarP(&mac.Port, "port", "P", 22, "Port.")
	command.PersistentFlags().StringVar(&mac.PrivateKey, "private_key", "", "PrivateKey.")
	command.PersistentFlags().StringVar(&mac.PublicKey, "public_key", "", "PublicKey.")
	command.PersistentFlags().StringVarP(&mac.Type, "type", "t", "password", "Type.")
	return command
}

func MachineRemoveCommand() *cobra.Command {
	var mac msgData.Machine
	var command = &cobra.Command{
		Use:   "rm machine_name",
		Short: "Remove machine.",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			mac.Name = args[0]
			mac.Del()
		},
	}
	return command
}

func MachineUpdateCommand() *cobra.Command {
	var mac msgData.Machine
	var command = &cobra.Command{
		Use:   "update machine_name",
		Short: "Update machine.",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			mac.Name = args[0]
			mac.Update()
		},
	}
	command.PersistentFlags().StringVarP(&mac.Host, "host", "H", "", "Host.")
	command.PersistentFlags().StringVarP(&mac.User, "user", "u", "root", "User.")
	command.PersistentFlags().StringVarP(&mac.Password, "password", "p", "123456", "Password.")
	command.PersistentFlags().IntVarP(&mac.Port, "port", "P", 22, "Port.")
	command.PersistentFlags().StringVar(&mac.PrivateKey, "private_key", "", "PrivateKey.")
	command.PersistentFlags().StringVar(&mac.PublicKey, "public_key", "", "PublicKey.")
	command.PersistentFlags().StringVarP(&mac.Type, "type", "t", "password", "Type.")
	return command
}

func MachineListCommand() *cobra.Command {
	var mac msgData.Machine
	var command = &cobra.Command{
		Use:   "ls",
		Short: "List machine.",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			mac.List()
		},
	}
	return command
}
