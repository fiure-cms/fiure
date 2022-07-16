package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fiure",
	Short: "FiureCMS cli",
	Long:  `Simple dstatic (dynamic and static website) generator cli tools for fiure cms`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// cobra auto init
	//cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fiure.yaml)")
	//rootCmd.MarkPersistentFlagRequired("config")

	//rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")

	// Viper flags lines
	//viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {

	// check config file
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// TODO: log lines add
		fmt.Println("Config file not found!")
		os.Exit(1)
	}

	// or look up automaticly
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			// TODO: log lines add
			fmt.Println("config file not found")
		} else {
			// Config file was found but another error was produced
			// TODO: log lines add
			fmt.Println("config file found but something wrong")
		}

		// TODO: log lines add
		fmt.Println("config file reader error:", err)
	} else {
		// TODO: log lines add
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func initCmdConfig(cmd string) {

	viper.SetConfigName(cmd)
	viper.SetConfigType("json")
	viper.AddConfigPath("configs")
	//viper.AllowEmptyEnv(true)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("dosya yok laaa")
			//viper.Set("list", map[string]interface{}{"obaraks": "ne dedin", "deneme": false})
			//viper.SafeWriteConfig()
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())
}
