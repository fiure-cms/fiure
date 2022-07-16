package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configBackupCmd represents the configBackup command
var configBackupCmd = &cobra.Command{
	Use:   "backup",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configBackup called")
	},
}

func init() {
	configCmd.AddCommand(configBackupCmd)
}
