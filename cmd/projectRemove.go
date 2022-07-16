package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// projectRemoveCmd represents the projectRemove command
var projectRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove selected project",
	Long:  `delete all configration files of selected project`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("projectRemove called")

		return nil
	},
}

func init() {
	projectCmd.AddCommand(projectRemoveCmd)
}
