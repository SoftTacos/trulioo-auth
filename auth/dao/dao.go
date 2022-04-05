package dao

//go:generate mockgen -package controller -destination mocks/mock_AuthDao.go . AuthDao

import (
	"errors"
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
	GetPassword(uuid string) (passwordHash string, err error)
}

type authDao struct {
	db *gopg.DB
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

func (d *authDao) GetPassword(uuid string) (passwordHash string, err error) {
	rows, err := d.db.Model(&passwordHash).Query("SELECT password_hash FROM passwords WHERE user_uuid = ?", uuid)
	if err != nil {
		log.Println("failed to retrieve password: ", err)
		return
	}

	if rows.RowsReturned() == 0 {
		err = errors.New("TODO: no rows")
		log.Println(err)
		return
	}

	return
}

type password struct {
	UserUuid     string
	PasswordHash string
	CreatedAt    *time.Time
	DeletedAt    *time.Time
}
