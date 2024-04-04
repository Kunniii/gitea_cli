package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/kunniii/gitea_cli/models"
	"golang.org/x/term"
)

func checkArgs(command string, action string, options ...string) bool {
	return true
}

func openInBrowser(url string) {
	cmd := exec.Command("open", url)
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting browser: ", err)
	}
}

func printHelp() {
	fmt.Println(lipgloss.NewStyle().
		Foreground(lipgloss.Color("10")).
		Bold(true).
		Render(
			"\nUsage: gitea <type> <status>\n\n" +
				"\t+ type  : issue, pull, branch\n" +
				"\t+ status: all, open, closed\n\n" +
				"Commands:\n" +
				"\t$ gitea issue open\n" +
				"\t$ gitea pull all\n" +
				"\t$ gitea branch\n"),
	)
}

func printError(err string) {
	var border = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("9")).
		Width(80).
		Padding(1, 2)
	errStyled := lipgloss.NewStyle().
		Foreground(lipgloss.Color("9")).
		Bold(true).
		Render(err)
	fmt.Println(border.Render(errStyled))
	os.Exit(1)
}

func printInfo(msg string) {
	var border = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		Width(80).
		Padding(1, 2)
	fmt.Println(border.Render(msg))
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

	var issueStateStyle string

	switch issue.State {
	case "open":
		issueStateStyle = issueTitleStyle.
			Foreground(lipgloss.Color("9")).Render("open")
	case "closed":
		issueStateStyle = issueTitleStyle.
			Foreground(lipgloss.Color("13")).Render("closed")
	}

	fmt.Println(border.Render(
		fmt.Sprintf(
			"%s [%s] %s\n%s\n-%s-",
			issueNumberString,
			issueStateStyle,
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
			printError(err.Error())
		}
		fmt.Println("\nYour token is saved at " + keyFilePath + "\n")
		return key
	}

	key, _ := os.ReadFile(keyFilePath)
	return strings.TrimSpace(string(key))
}
