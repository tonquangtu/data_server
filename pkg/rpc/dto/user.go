package dto

import "time"

type AddUserDto struct {
	Username  string
	Password  string
	CreatedAt time.Time
}

type GetUserDto struct {
	ID           uint
	Username     string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
