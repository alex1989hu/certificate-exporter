//go:build generatedocs

package main

import (
	"log"

	"github.com/spf13/cobra/doc"

	"github.com/patrickmsieber/certificate-exporter/cmd"
)

func main() {
	log.Println("Generate Markdown Documentation for CLI commands")

	if err := doc.GenMarkdownTree(cmd.GetRootCmd(), "./docs"); err != nil {
		log.Fatal(err)
	}

	log.Println("Markdown Documentation Generation is done")
}
