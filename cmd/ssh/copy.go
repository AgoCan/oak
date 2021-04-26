package ssh

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Recursion bool

	CopyCmd = &cobra.Command{
		Use:   "cp [machine]",
		Short: "copy file or director.",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}
)

func init() {
	CopyCmd.PersistentFlags().BoolVarP(&Recursion, "recursion", "r", false, "Recursion director.")
}
