package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// bindFlags binds each cobra flag to its associated viper configuration (e.g. environment variable).
func bindFlags(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && viper.IsSet(f.Name) {
			if err := cmd.Flags().Set(f.Name, viper.GetString(f.Name)); err != nil {
				cobra.CheckErr(err)
			}
		}
	})
}
