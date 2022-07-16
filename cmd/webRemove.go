package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// webRemoveCmd represents the webRemove command
var webRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("webRemove called")
	},
}

func init() {
	webCmd.AddCommand(webRemoveCmd)
}
