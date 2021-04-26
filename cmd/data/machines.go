package data

import (
	msgData "github.com/agocan/oak/pkg/data"
	"github.com/spf13/cobra"
)

var (
	MacCmd = &cobra.Command{
		Use:   "mac [command]",
		Short: "Operating machine.",
		Long:  ``,
	}
	MacAddCmd = &cobra.Command{
		Use:   "add machine_name",
		Short: "Add machine.",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run:   AddMachine,
	}
	MacRemoveCmd = &cobra.Command{
		Use:   "rm machine_name",
		Short: "Remove machine.",
		Long:  ``,
		Run:   RemoveMachine,
	}
	MacUpdateCmd = &cobra.Command{
		Use:   "update machine_name",
		Short: "Update machine.",
		Long:  ``,
		Run:   UpdateMachine,
	}
	MacListCmd = &cobra.Command{
		Use:   "ls",
		Short: "List machine.",
		Long:  ``,
		Run:   ListMachine,
	}
	mac msgData.Machine
)

func init() {
	MacCmd.AddCommand(MacAddCmd)
	MacCmd.AddCommand(MacRemoveCmd)
	MacCmd.AddCommand(MacUpdateCmd)
	MacCmd.AddCommand(MacListCmd)

	MacAddCmd.PersistentFlags().StringVarP(&mac.Host, "host", "H", "", "Host.")
	MacAddCmd.PersistentFlags().StringVarP(&mac.User, "user", "u", "root", "User.")
	MacAddCmd.PersistentFlags().StringVarP(&mac.Password, "password", "p", "123456", "Password.")
	MacAddCmd.PersistentFlags().IntVarP(&mac.Port, "port", "P", 22, "Port.")
	MacAddCmd.PersistentFlags().StringVar(&mac.PrivateKey, "private_key", "", "PrivateKey.")
	MacAddCmd.PersistentFlags().StringVar(&mac.PublicKey, "public_key", "", "PublicKey.")
	MacAddCmd.PersistentFlags().StringVarP(&mac.Type, "type", "t", "password", "Type.")

	MacUpdateCmd.PersistentFlags().StringVarP(&mac.Host, "host", "H", "", "Host.")
	MacUpdateCmd.PersistentFlags().StringVarP(&mac.User, "user", "u", "root", "User.")
	MacUpdateCmd.PersistentFlags().StringVarP(&mac.Password, "password", "p", "123456", "Password.")
	MacUpdateCmd.PersistentFlags().IntVarP(&mac.Port, "port", "P", 22, "Port.")
	MacUpdateCmd.PersistentFlags().StringVar(&mac.PrivateKey, "private_key", "", "PrivateKey.")
	MacUpdateCmd.PersistentFlags().StringVar(&mac.PublicKey, "public_key", "", "PublicKey.")
	MacUpdateCmd.PersistentFlags().StringVarP(&mac.Type, "type", "t", "password", "Type.")
}

func AddMachine(cmd *cobra.Command, args []string) {
	mac.Name = args[0]
	mac.Add()
}

func RemoveMachine(cmd *cobra.Command, args []string) {
	mac.Name = args[0]
	mac.Del()
}

func UpdateMachine(cmd *cobra.Command, args []string) {
	mac.Name = args[0]
	mac.Update()
}

func ListMachine(cmd *cobra.Command, args []string) {
	mac.List()
}
