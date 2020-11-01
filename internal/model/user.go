package model

import (
	"time"
)

type User struct {
	ID           uint
	Username     string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

func (User) TableName() string {
	return "user"
}
