package model

import (
	"time"
)

type User struct {
	ID             int
	Email          string
	Username       string
	Password       string
	HashedPassword string
}

type Session struct {
	ID      int
	UserID  int
	UUID    string
	ExpDate time.Time
}
