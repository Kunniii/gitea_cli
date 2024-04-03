package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/kunniii/gitea_cli/models"
)

type Gitea struct {
	Token      string
	BaseURL    string
	HTTPClient *http.Client
}

func NewGitea() *Gitea {
	client := &http.Client{}
	return &Gitea{HTTPClient: client}
}

func (gt *Gitea) WithURL(urlString string) *Gitea {
	u, err := url.Parse(urlString)
	if err != nil {
		log.Fatal("Cannot parse URL: ", urlString)
	}
	reqURI := u.RequestURI()
	reqURI, _ = strings.CutSuffix(reqURI, ".git")
	gt.BaseURL = "https://" + u.Hostname() + "/api/v1/repos" + reqURI
	return gt
}

func (gt *Gitea) WithToken(token string) *Gitea {
	gt.Token = token
	return gt
}

func (gt *Gitea) getIssues(state string, issueType string) ([]models.Issue_Pull, error) {
	if state == "" {
		state = "open"
	}
	if issueType == "" {
		issueType = "issues"
	}

	requestURL := gt.BaseURL + "/issues/" + "?state=" + state + "&type=" + issueType

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "token "+gt.Token)
	resp, err := gt.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var issues []models.Issue_Pull
	err = json.Unmarshal(body, &issues)
	if err != nil {
		return nil, err
	}
	return issues, nil
}
