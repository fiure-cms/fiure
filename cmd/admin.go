package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// adminCmd represents the admin command
var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "list of admin panel",
	Long:  `list all of the admin panels at once.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("admin called")
	},
}

func init() {
	rootCmd.AddCommand(adminCmd)
}
