package model

type User struct {
	ID             int
	Email          string
	Username       string
	Password       string
	HashedPassword string
}
