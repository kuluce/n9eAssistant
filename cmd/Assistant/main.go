package main

import (
	"n9eHelper/internal/command"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "n9eHelper",
		Short: "n9eHelper is a tool for n9e",
		Long:  `n9eHelper is a tool for n9e`,
		Example: ` 
	n9eHelper version
	n9eHelper help
	`,
	}
)

func main() {
	// logger.Infof("Hello, World!")
	// logger.Errorf("Error message!")

	// logger.Warn("warn message")
	// logger.Infof("end")
	rootCmd.AddCommand(command.VersionCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}
