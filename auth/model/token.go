package model

import "time"

type RefreshToken struct {
	UserUuid  string
	Header    string
	Payload   string
	Signature string
	ExpiresAt *time.Time
}
