package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// projectAddCmd represents the projectAdd command
var projectAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add new project",
	Long:  `create new project configration`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("projectAdd called")

		viper.Set("list", map[string]interface{}{"obaraks": "ne dedin", "deneme": false})
		return viper.WriteConfig()
	},
}

func init() {
	projectCmd.AddCommand(projectAddCmd)
}
