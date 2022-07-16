package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// buildAllCmd represents the buildAll command
var buildAllCmd = &cobra.Command{
	Use:   "all",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("buildAll called")
	},
}

func init() {
	buildCmd.AddCommand(buildAllCmd)
}
