package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	category string
	action   string
	options  []string
)

func init() {
	// gt issue    list       all
	// gt <category> <action> <option>
	// gt branch   list
	// gt issue view id
	// gt pull view id
	// gt repo view --web

	if len(os.Args) == 1 {
		printHelp()
		os.Exit(0)
	}

	category = os.Args[1]

	if len(os.Args) > 2 {
		action = os.Args[2]
	}
	if len(os.Args) > 3 {
		options = os.Args[3:]
	}

}

func main() {
	var token = loadToken()
	var err error

	urlString, err := getURL()
	if err != nil {
		printError(err.Error())
	}

	urlString, _ = strings.CutSuffix(urlString, "\n")

	var gt = NewGitea().
		WithToken(token).
		WithURL(urlString)

	switch category {
	case "branch":
		branches, err := gt.GetBranches()
		if err != nil {
			printError(err.Error())
		}
		prettyPrintBranches(branches)
	case "issue", "pull":
		results, err := gt.GetIssues(category, action)
		if err != nil {
			printError(err.Error())
		}
		if len(results) == 0 {
			printInfo(fmt.Sprintf("There are %s no %ss on this repository", action, category))
		}
		for _, data := range results {
			prettyPrintIssue_Pull(data)
		}
	}
}
