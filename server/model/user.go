package model

type User struct {
	Id       string
	Password string
	Uuid     string
	Tasks    []Task
}
