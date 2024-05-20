package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// simple api rest services
var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "MyApp is a CLI application",
	Long:  `MyApp is a CLI application that demonstrates how to build a simple REST API service`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
