package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/patrickmsieber/certificate-exporter/version"
)

//nolint:gochecknoglobals
var (
	// isDebug represents the debug level of logger.
	isDebug bool

	// rootCmd represents the base command when called without any subcommands.
	rootCmd = &cobra.Command{
		DisableAutoGenTag: true,
		Use:               "certificate-exporter",
		Version:           fmt.Sprintf("(%s/%s)", version.GitBranch, version.GitCommit),
		Short:             "CLI to export certificates in PEM format",
	}
)

// GetRootCmd returns root cobra command to let us able to generate CLI documentation.
func GetRootCmd() *cobra.Command {
	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Execution failed: %v", err) //nolint:revive // We must immediately exit, deep-exit.
	}
}

func init() {
	viper.SetEnvPrefix("CERTIFICATE_EXPORTER")

	viper.AutomaticEnv()

	// Match dashes with underscores in environment variables
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVarP(&isDebug, "debug", "d", false,
		"enable debug logging")
}

// initConfig called by cobra.
func initConfig() {
	// Connect persistent flags and environment variables
	if err := viper.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		cobra.CheckErr(err)
	}
}
