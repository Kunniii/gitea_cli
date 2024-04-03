package main

import (
	"log"
	"os"
	"strings"
)

var (
	status    string
	issueType string
)

func init() {
	if len(os.Args) == 1 {
		printHelp()
		os.Exit(0)
	}

	status = os.Args[1]
	if len(os.Args) > 2 {
		issueType = os.Args[2]
	}

	if status != "" && status != "branch" && status != "all" && status != "open" && status != "closed" {
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

	if status == "branch" {
		branches, err := gt.getBranches()
		if err != nil {
			log.Fatal(err)
		}
		prettyPrintBranches(branches)
	} else {
		results, err := gt.getIssues(status, issueType)
		if err != nil {
			log.Fatal(err)
		}
		for _, data := range results {
			prettyPrintIssue_Pull(data)
		}

	}

}
