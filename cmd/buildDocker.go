package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// buildDockerCmd represents the buildDocker command
var buildDockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("buildDocker called")
	},
}

func init() {
	buildCmd.AddCommand(buildDockerCmd)
}
