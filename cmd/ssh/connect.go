package ssh

import (
	"fmt"

	"github.com/spf13/cobra"
)

var TestConnectCmd = &cobra.Command{
	Use:   "test",
	Short: "Test connect machine.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}
