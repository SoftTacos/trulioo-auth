package model

import (
	"time"
)

type User struct {
	UUID      string
	Email     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	// PhoneNumber string
	// FirstName   string
	// LastName    string
}
