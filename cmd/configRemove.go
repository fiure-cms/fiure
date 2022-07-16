package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configRemoveCmd represents the configRemove command
var configRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configRemove called")
	},
}

func init() {
	configCmd.AddCommand(configRemoveCmd)
}
