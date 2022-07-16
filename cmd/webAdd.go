package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// webAddCmd represents the webAdd command
var webAddCmd = &cobra.Command{
	Use:   "add",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("webAdd called")
	},
}

func init() {
	webCmd.AddCommand(webAddCmd)
}
