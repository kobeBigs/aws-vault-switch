package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aws-vault-switch",
	Short: "A simple CLI app that extends aws-vault to switch between profiles",
	Long:  `aws-vault-switch is a CLI tool designed to enable seamless switching between AWS profiles without the need to exit the current session`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
