package model

import "time"

type Task struct {
	Id        int    `json:"id"`
	Content   string `json:"content"`
	Deadline  string `json:"deadline"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      *User
}
