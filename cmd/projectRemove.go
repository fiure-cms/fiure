package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// projectRemoveCmd represents the projectRemove command
var projectRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove selected project",
	Long:  `delete all configration files of selected project`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("projectRemove called")

		return Unset("list.obaraks")
	},
}

func init() {
	projectCmd.AddCommand(projectRemoveCmd)
}

// Via: https://github.com/spf13/viper/issues/632#issuecomment-869668629
func Unset(vars ...string) error {
	cfg := viper.AllSettings()
	vals := cfg

	for _, v := range vars {
		parts := strings.Split(v, ".")
		for i, k := range parts {
			v, ok := vals[k]
			if !ok {
				// Doesn't exist no action needed
				break
			}

			switch len(parts) {
			case i + 1:
				// Last part so delete.
				delete(vals, k)
			default:
				m, ok := v.(map[string]interface{})
				if !ok {
					return fmt.Errorf("unsupported type: %T for %q", v, strings.Join(parts[0:i], "."))
				}
				vals = m
			}
		}
	}

	b, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		return err
	}

	if err = viper.ReadConfig(bytes.NewReader(b)); err != nil {
		return err
	}

	return viper.WriteConfig()
}
