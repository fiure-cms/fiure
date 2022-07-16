package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deployLocalCmd represents the deployLocal command
var deployLocalCmd = &cobra.Command{
	Use:   "local",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deployLocal called")
	},
}

func init() {
	deployCmd.AddCommand(deployLocalCmd)

}
