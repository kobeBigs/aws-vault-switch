package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

func listAndSwitchProfiles() {
	cmd := exec.Command("aws-vault", "list")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error listing profiles: %v\n", err)
		return
	}

	profilesList := strings.Split(strings.TrimSpace(string(output)), "\n")

	var profiles []string
	for _, profileItem := range profilesList[2:] {
		profileItems := strings.Fields(profileItem)
		if len(profileItems) > 0 {
			profiles = append(profiles, profileItems[0])
		}
	}
	if len(profiles) == 0 {
		fmt.Println("No profiles found.")
		return
	}

	prompt := promptui.Select{
		Label: "Select AWS Profile",
		Items: profiles,
	}
	_, selectedProfile, err := prompt.Run()
	if err != nil {
		fmt.Printf("Selection failed: %v\n", err)
		return
	}

	switchProfile(selectedProfile)
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
