package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// projectSelectCmd represents the projectSelect command
var projectSelectCmd = &cobra.Command{
	Use:   "select",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("projectSelect called")

		viper.Set("current", "my")

		return viper.WriteConfig()
	},
}

func init() {
	projectCmd.AddCommand(projectSelectCmd)
}
