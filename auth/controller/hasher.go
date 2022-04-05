package controller

//go:generate mockgen -package controller -destination mocks/mock_PasswordHasher.go . PasswordHasher

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher interface {
	HashAndSalt(password []byte, cost int) (hash []byte, err error)
}

type pwHasher struct{}

func (h pwHasher) HashAndSalt(password []byte, cost int) (hash []byte, err error) {
	hash, err = bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return
}
