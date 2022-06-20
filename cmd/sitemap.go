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
	Run:   fiureSitemap,
}

func init() {
	rootCmd.AddCommand(sitemapCmd)
	sitemapCmd.AddCommand(sitemapCreateCmd, sitemapClearCmd)
}

// sitemapCreateCmd represents the sitemap command
var sitemapCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new sitemap",
	Long:  `Generate new sitemap index files with sub files before clear target directory`,
	RunE:  createSitemap,
}

// sitemapClearCmd represents the sitemap command
var sitemapClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear old sitemap files",
	Long:  `Clear target directory`,
	RunE:  clearSitemap,
}

// sitemap main and sub commands function
func fiureSitemap(cmd *cobra.Command, args []string) {
	fmt.Println("sitemap called")
}

func createSitemap(cmd *cobra.Command, args []string) error {
	fmt.Println("sitemapCreate called")

	return nil
}

func clearSitemap(cmd *cobra.Command, args []string) error {
	fmt.Println("sitemapClear called")

	return nil
}
