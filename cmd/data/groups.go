package data

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "group [command]",
	Short: "Operating group.",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run:   GroupOperate,
}

func GroupOperate(cmd *cobra.Command, args []string) {
}
