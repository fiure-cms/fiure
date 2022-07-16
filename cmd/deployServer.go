package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deployServerCmd represents the deployServer command
var deployServerCmd = &cobra.Command{
	Use:   "server",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deployServer called")
	},
}

func init() {
	deployCmd.AddCommand(deployServerCmd)
}
