package models

import "time"

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
