package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var profile string

var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch aws profiles without exiting",
	Long:  `Use switch to change aws profiles without exiting from current shell session `,
	Run: func(cmd *cobra.Command, args []string) {
		switchProfile(profile)
	},
}

func switchProfile(profile string) {
	os.Unsetenv("AWS_VAULT")
	cmd := exec.Command("aws-vault", "exec", profile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error switching profile: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(switchCmd)
	switchCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS profile to switch to")
	switchCmd.MarkFlagRequired("profile")
}
