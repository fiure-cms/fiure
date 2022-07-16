package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// adminAddCmd represents the adminAdd command
var adminAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add new admin panel",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("adminAdd called")
	},
}

func init() {
	adminCmd.AddCommand(adminAddCmd)
}
