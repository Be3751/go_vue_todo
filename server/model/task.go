package model

import "time"

type Task struct {
	Id        int    `json:"id"`
	Content   string `json:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      *User
}
