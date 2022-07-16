package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// projectInstallCmd represents the projectInstall command
var projectInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("projectInstall called")

		return nil
	},
}

func init() {
	projectCmd.AddCommand(projectInstallCmd)
}
