package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// buildDockerComposeCmd represents the buildDockerCompose command
var buildDockerComposeCmd = &cobra.Command{
	Use:   "docker-compose",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("buildDockerCompose called")
	},
}

func init() {
	buildCmd.AddCommand(buildDockerComposeCmd)
}
