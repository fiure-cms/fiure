package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// adminRemoveCmd represents the adminRemove command
var adminRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove selected admin panel",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("adminRemove called")
	},
}

func init() {
	adminCmd.AddCommand(adminRemoveCmd)
}
