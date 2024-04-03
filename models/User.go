package models

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
