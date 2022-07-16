package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// adminSelectCmd represents the adminSelect command
var adminSelectCmd = &cobra.Command{
	Use:   "select",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("adminSelect called")
	},
}

func init() {
	adminCmd.AddCommand(adminSelectCmd)
}
