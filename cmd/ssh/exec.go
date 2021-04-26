package ssh

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ExecCmd = &cobra.Command{
	Use:   "exec [machine]",
	Short: "Excute command.",
	Long:  ``,
	Run:   Exec,
}

func Exec(cmd *cobra.Command, args []string) {
	fmt.Println(args)
}
