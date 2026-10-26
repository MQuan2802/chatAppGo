package entity

import (
	"time"
)

type User struct {
	Id           uint
	Phone        string
	Username     string
	PasswordHash string
	Mail         string
	CreatedAt    time.Time
	LastUpdateAt time.Time
}
