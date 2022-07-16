package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// projectBackupCmd represents the projectBackup command
var projectBackupCmd = &cobra.Command{
	Use:   "backup",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("projectBackup called")

		return nil
	},
}

func init() {
	projectCmd.AddCommand(projectBackupCmd)
}
