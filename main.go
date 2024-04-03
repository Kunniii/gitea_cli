package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/kunniii/gitea_cli/models"
)

func prettyPrint(issue models.Issue_Pull) {
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

var (
	status    string
	issueType string
)

func init() {
	// check os.args
	if len(os.Args) == 1 {
		status = ""
		issueType = ""
	} else if len(os.Args) > 1 {
		status = os.Args[1]
		issueType = ""
	} else if len(os.Args) > 2 {
		status = os.Args[1]
		issueType = os.Args[2]
	}

	if status != "" && status != "all" && status != "open" && status != "closed" {
		printHelp()
		os.Exit(0)
	}

	if issueType != "" && issueType != "issues" && issueType != "pulls" {
		printHelp()
		os.Exit(0)
	}
}

func main() {
	var token = loadToken()
	var err error

	urlString, err := getURL()
	if err != nil {
		log.Fatal(err)
	}

	urlString, _ = strings.CutSuffix(urlString, "\n")

	var gt = NewGitea().
		WithToken(token).
		WithURL(urlString)

	results, err := gt.getIssues(status, issueType)
	if err != nil {
		log.Fatal(err)
	}

	for _, issue := range results {
		prettyPrint(issue)
	}

}
