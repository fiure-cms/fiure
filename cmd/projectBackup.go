package cmd

import (
	"fmt"

	"github.com/fikir-uretgeci/fiure/tools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// projectBackupCmd represents the projectBackup command
var projectBackupCmd = &cobra.Command{
	Use:   "backup",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("projectBackup called")

		path := "./backup/configs/"
		err := tools.CreateDir(path)
		if err != nil {
			return err
		}

		return viper.WriteConfigAs(fmt.Sprintf("%s%s", path, "projects.bck.json"))
	},
}

func init() {
	projectCmd.AddCommand(projectBackupCmd)
}
