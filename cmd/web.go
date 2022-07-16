package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("web called")
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
}
