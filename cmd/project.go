package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "list of projects",
	Long:  `list all of the projects at once.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("project called")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
}
