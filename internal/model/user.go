package model

import (
	"time"

	"github.com/google/uuid"
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
	UUID    uuid.UUID
	ExpDate time.Time
}
