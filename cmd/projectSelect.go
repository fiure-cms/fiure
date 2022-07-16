package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// projectSelectCmd represents the projectSelect command
var projectSelectCmd = &cobra.Command{
	Use:   "select",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("projectSelect called")

		return nil
	},
}

func init() {
	projectCmd.AddCommand(projectSelectCmd)
}
