package data

import (
	"github.com/spf13/cobra"

	msgData "github.com/agocan/oak/pkg/data"
)

func GroupCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "group [command]",
		Short: "Operating groups.",
		Long:  ``,
	}
	command.AddCommand(GroupAddCommand())
	command.AddCommand(GroupRemoveCommand())
	command.AddCommand(GroupUpdateCommand())
	command.AddCommand(GroupListCommand())

	command.AddCommand(GroupMachineCommand())

	return command
}

func GroupAddCommand() *cobra.Command {
	var group msgData.Group
	var command = &cobra.Command{
		Use:   "add group name",
		Short: "Add group.",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			group.Name = args[0]
			group.Add()
		},
	}
	return command
}

func GroupRemoveCommand() *cobra.Command {
	var group msgData.Group
	var command = &cobra.Command{
		Use:   "rm group name",
		Short: "Remove group.",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			group.Name = args[0]
			group.Del()
		},
	}
	return command
}

func GroupUpdateCommand() *cobra.Command {
	var group msgData.Group
	var command = &cobra.Command{
		Use:   "update group_name new_group_name",
		Short: "Update group.",
		Long:  ``,
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			group.Name = args[0]

			group.Update(args[1])
		},
	}
	return command
}

func GroupListCommand() *cobra.Command {
	var group msgData.Group
	var command = &cobra.Command{
		Use:   "ls",
		Short: "List group.",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			group.List()
		},
	}
	return command
}
