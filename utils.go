package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/kunniii/gitea_cli/models"
	"golang.org/x/term"
)

func printHelp() {
	fmt.Println(lipgloss.NewStyle().
		Foreground(lipgloss.Color("10")).
		Bold(true).
		Render(
			"\nUsage: gitea <status> <type>\n\n" +
				"\t+ status: all, open, closed\n" +
				"\t+ type  : issues, pulls\n\n" +
				"Commands:\n" +
				"\t$ gitea all issues\n" +
				"\t$ gitea open pulls\n"),
	)
}

func prettyPrintBranches(branches []models.Branch) {
	var border = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		Width(80).
		Padding(0, 2)
	var branchStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("0"))

	var data = "\n"
	for _, branch := range branches {
		data += branchStyle.Render(branch.Name) + "\n"
	}
	fmt.Println(border.Render(data))
}

func prettyPrintIssue_Pull(issue models.Issue_Pull) {
	var border = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		Width(80).
		Padding(1, 2)

	var issueTitleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("10"))

	var issueNumberStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("12"))

	var issueNumberString = issueNumberStyle.Render(fmt.Sprintf("#%d", issue.Number))

	var issueStyleString = issueTitleStyle.Render(issue.Title)

	fmt.Println(border.Render(
		fmt.Sprintf(
			"%s: %s\n%s\n-%s-",
			issueNumberString,
			issueStyleString,
			issue.HTMLURL,
			issue.User.Login,
		),
	))

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
