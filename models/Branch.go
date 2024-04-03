package models

import "time"

type Branch struct {
	Name   string `json:"name"`
	Commit struct {
		ID      string `json:"id"`
		Message string `json:"message"`
		URL     string `json:"url"`
		Author  struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"author"`
		Committer struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"committer"`
		Verification struct {
			Verified  bool   `json:"verified"`
			Reason    string `json:"reason"`
			Signature string `json:"signature"`
			Signer    any    `json:"signer"`
			Payload   string `json:"payload"`
		} `json:"verification"`
		Timestamp time.Time `json:"timestamp"`
		Added     any       `json:"added"`
		Removed   any       `json:"removed"`
		Modified  any       `json:"modified"`
	} `json:"commit"`
	Protected                     bool   `json:"protected"`
	RequiredApprovals             int    `json:"required_approvals"`
	EnableStatusCheck             bool   `json:"enable_status_check"`
	StatusCheckContexts           []any  `json:"status_check_contexts"`
	UserCanPush                   bool   `json:"user_can_push"`
	UserCanMerge                  bool   `json:"user_can_merge"`
	EffectiveBranchProtectionName string `json:"effective_branch_protection_name"`
}
