package model

type User struct {
	Id       int
	Uuid     string
	Password string
}

type Session struct {
	Id   int
	Uuid string
}
