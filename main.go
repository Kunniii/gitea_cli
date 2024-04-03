package main

import (
	"log"
	"os"
	"strings"
)

var (
	status  string
	theType string
)

func init() {
	if len(os.Args) == 1 {
		printHelp()
		os.Exit(0)
	}

	theType = os.Args[1]
	if len(os.Args) > 2 {
		status = os.Args[2]
	}

	if theType != "" && theType != "issue" && theType != "pull" && theType != "branch" {
		printHelp()
		os.Exit(0)
	}

	if status != "" && status != "all" && status != "open" && status != "closed" {
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

	if theType == "branch" {
		branches, err := gt.getBranches()
		if err != nil {
			log.Fatal(err)
		}
		prettyPrintBranches(branches)
	} else {
		results, err := gt.getIssues(status, theType)
		if err != nil {
			log.Fatal(err)
		}
		for _, data := range results {
			prettyPrintIssue_Pull(data)
		}

	}

}
