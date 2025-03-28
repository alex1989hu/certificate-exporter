package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/patrickmsieber/certificate-exporter/pkg/exporter"
)

//nolint:gochecknoglobals
var (
	// exportCmd represents the generate command.
	exportCmd = &cobra.Command{
		Use:   "export",
		Short: "Reads certificate from a CSV file and export them separately",
		Long:  "",
		Run: func(_ *cobra.Command, _ []string) {
			start()
		},
	}

	// csvFilePath represents CSV file path which contains certificates from the database.
	csvFilePath string

	// outputDirectoryPath represents the directory path which will contain exported certificates.
	outputDirectoryPath string
)

func init() {
	rootCmd.AddCommand(exportCmd)

	cobra.OnInitialize(func() {
		// Connect flags and environment variables
		bindFlags(exportCmd)

		for _, flag := range []string{"csv-file"} {
			if err := exportCmd.MarkFlagRequired(flag); err != nil {
				cobra.CheckErr(err)
			}
		}
	})

	exportCmd.Flags().StringVarP(&csvFilePath, "csv-file", "c", "",
		"set csv file")

	exportCmd.Flags().StringVar(&outputDirectoryPath, "output-directory", "output",
		"set output directory")
}

func start() {
	log.Printf("Input CSV file: %s\n", csvFilePath)

	// A possible way to continue the coding challenge
	if err := exporter.Export(csvFilePath, outputDirectoryPath); err != nil {
		log.Fatal(err) //nolint:revive // We must immediately exit, deep-exit.
	}
}
