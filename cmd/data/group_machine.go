package data

import (
	msgData "github.com/agocan/oak/pkg/data"
	"github.com/spf13/cobra"
)

func GroupMachineCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "mac ",
		Short: "Group operate Machine.",
		Long:  ``,
	}
	command.AddCommand(GroupAddMachineCommand())
	command.AddCommand(GroupRemoveMachineCommand())
	command.AddCommand(GroupListMachineCommand())
	return command
}

func GroupAddMachineCommand() *cobra.Command {
	var group msgData.Group

	var command = &cobra.Command{
		Use:   "add group_name machine_name",
		Short: "Group add Machine.",
		Long:  ``,
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			group.Name = args[0]
			group = *msgData.GetGroup(group.Name)
			group.Machines = append(group.Machines, args[1:]...)
			group.AddMachine()
		},
	}
	return command
}

func GroupRemoveMachineCommand() *cobra.Command {
	var group msgData.Group
	var command = &cobra.Command{
		Use:   "rm group_name machine_name",
		Short: "Group remove Machine.",
		Long:  ``,
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			group.Name = args[0]
			group.DelMachine(args[1:])
		},
	}
	return command
}

func GroupListMachineCommand() *cobra.Command {
	var group msgData.Group
	var command = &cobra.Command{
		Use:   "ls group_name",
		Short: "Group list Machine.",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			group.Name = args[0]
			group.ListMachine()
		},
	}
	return command
}
