package dao

//go:generate mockgen -package controller -destination mocks/mock_AuthDao.go . AuthDao

import (
	"log"
	"time"

	gopg "github.com/go-pg/pg/v10"
)

func NewAuthDao(db *gopg.DB) AuthDao {
	return &authDao{
		db: db,
	}
}

type AuthDao interface {
	CreatePassword(uuid string, passwordHash string) (err error)
}

type authDao struct {
	db *gopg.DB
}

// this is in the dao because the rest of the service doesn't need to see the data model
type password struct {
	UserUuid     string
	PasswordHash string
	CreatedAt    *time.Time
	DeletedAt    *time.Time
}

func (d *authDao) CreatePassword(uuid string, passwordHash string) (err error) {
	now := time.Now()
	pw := password{
		UserUuid:     uuid,
		PasswordHash: passwordHash,
		CreatedAt:    &now,
	}

	_, err = d.db.Model(&pw).Insert()
	if err != nil {
		log.Println("failed to create password: ", err)
	}
	return
}

func (d *authDao) CreateRefreshToken(uuid string, passwordHash string) (err error) {
	now := time.Now()
	pw := password{
		UserUuid:     uuid,
		PasswordHash: passwordHash,
		CreatedAt:    &now,
	}

	_, err = d.db.Model(&pw).Insert()
	if err != nil {
		log.Println("failed to create password: ", err)
	}
	return
}
