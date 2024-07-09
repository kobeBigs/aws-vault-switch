package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aws-vault-switch",
	Short: "Switch between AWS profiles",
	Long:  `A simple CLI app that extends aws-vault to allow switching between managed AWS profiles`,
	Run: func(cmd *cobra.Command, args []string) {
		listAndSwitchProfiles()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
