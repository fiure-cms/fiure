package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sitemapCmd represents the sitemap command
var sitemapCmd = &cobra.Command{
	Use:   "sitemap",
	Short: "Sitemap generator",
	Long:  `Commands of sitemap generator`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sitemap called")
	},
}

func init() {
	sitemapCmd.AddCommand(sitemapCreateCmd, sitemapClearCmd)
	rootCmd.AddCommand(sitemapCmd)
}

// sitemapCreateCmd represents the sitemap command
var sitemapCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new sitemap",
	Long:  `Generate new sitemap index files with sub files before clear target directory`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sitemapCreate called")
	},
}

// sitemapClearCmd represents the sitemap command
var sitemapClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear old sitemap files",
	Long:  `Clear target directory`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sitemapClear called")
	},
}
