package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

func printHelp() {
	fmt.Println(lipgloss.NewStyle().
		Foreground(lipgloss.Color("10")).
		Bold(true).
		Render(
			"\nUsage: gitea <status> <type>\n\n" +
				"Commands:\n" +
				"\t$ gitea all issues\n" +
				"\t$ gitea open pulls\n"),
	)
}

func getURL() (string, error) {
	cmd := exec.Command("git", "remote", "get-url", "origin")
	stdout, err := cmd.Output()
	return string(stdout), err
}

func loadToken() string {
	homeDir, _ := os.UserHomeDir()
	keyFilePath := homeDir + "/.gitea_cli_token"

	if _, err := os.Stat(keyFilePath); os.IsNotExist(err) {
		fmt.Print("Please enter your Token: ")
		bytePassword, _ := term.ReadPassword(0)
		key := string(bytePassword)
		if err := os.WriteFile(keyFilePath, []byte(key), 0600); err != nil {
			log.Fatal(err)
		}
		fmt.Println("\nYour token is saved at " + keyFilePath + "\n")
		return key
	}

	key, _ := os.ReadFile(keyFilePath)
	return strings.TrimSpace(string(key))
}
