package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type User struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
	Language  string `json:"language"`
	IsAdmin   bool   `json:"is_admin"`
	LastLogin string `json:"last_login"`
	Created   string `json:"created"`
	Username  string `json:"username"`
}

type Repository struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	FullName string `json:"full_name"`
}

type Issue struct {
	ID               int         `json:"id"`
	URL              string      `json:"url"`
	HTMLURL          string      `json:"html_url"`
	Number           int         `json:"number"`
	User             User        `json:"user"`
	OriginalAuthor   string      `json:"original_author"`
	OriginalAuthorID int         `json:"original_author_id"`
	Title            string      `json:"title"`
	Body             string      `json:"body"`
	Labels           []string    `json:"labels"`
	Milestone        interface{} `json:"milestone"`
	Assignee         interface{} `json:"assignee"`
	Assignees        interface{} `json:"assignees"`
	State            string      `json:"state"`
	IsLocked         bool        `json:"is_locked"`
	Comments         int         `json:"comments"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	ClosedAt         interface{} `json:"closed_at"`
	DueDate          interface{} `json:"due_date"`
	PullRequest      interface{} `json:"pull_request"`
	Repository       Repository  `json:"repository"`
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

func prettyPrintIssue(issue Issue) {
	var border = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		Padding(1, 2)

	var issueTitleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("47"))

	var issueNumberStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("63"))

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

func main() {

	var token = loadToken()

	var urlString, _ = getURL()
	var url, _ = url.Parse(urlString)
	var reqURI = url.RequestURI()
	reqURI, _ = strings.CutSuffix(reqURI, ".git")
	var apiURL = "https://" + url.Hostname() + "/api/v1/repos" + reqURI + "/issues?state=open&type=issues"

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "token "+token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var issues []Issue
	err = json.Unmarshal(body, &issues)
	if err != nil {
		log.Fatal(err)
	}

	for _, issue := range issues {
		prettyPrintIssue(issue)
	}

}
