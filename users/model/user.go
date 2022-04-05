package model

import (
	"time"

	v1 "github.com/softtacos/trulioo-auth/grpc/users/v1"
)

type User struct {
	*v1.User
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
