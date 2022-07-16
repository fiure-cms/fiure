package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configAddCmd represents the configAdd command
var configAddCmd = &cobra.Command{
	Use:   "add",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configAdd called")
	},
}

func init() {
	configCmd.AddCommand(configAddCmd)
}
