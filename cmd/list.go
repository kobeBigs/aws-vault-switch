package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists aws profiles",
	Long:  `Lists profiles managed by aws-vault.`,
	Run: func(cmd *cobra.Command, args []string) {
		listProfiles()
	},
}

func listProfiles() {
	cmd := exec.Command("aws-vault", "list")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error listing profiles: %v\n", err)
		return
	}

	profiles := strings.TrimSpace(string(output))
	if profiles == "" {
		fmt.Println("No profiles found.")
	} else {
		fmt.Println("AWS Vault Profiles:")
		fmt.Println(profiles)
	}
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
