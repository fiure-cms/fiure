package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// projectAddCmd represents the projectAdd command
var projectAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add new project",
	Long:  `create new project configration`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("projectAdd called")

		return nil
	},
}

func init() {
	projectCmd.AddCommand(projectAddCmd)
}
