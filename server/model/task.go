package model

type Task struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	User    *User
}
