package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// buildGoCmd represents the buildGo command
var buildGoCmd = &cobra.Command{
	Use:   "go",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("buildGo called")
	},
}

func init() {
	buildCmd.AddCommand(buildGoCmd)
}
