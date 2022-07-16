package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// webSelectCmd represents the webSelect command
var webSelectCmd = &cobra.Command{
	Use:   "select",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("webSelect called")
	},
}

func init() {
	webCmd.AddCommand(webSelectCmd)
}
