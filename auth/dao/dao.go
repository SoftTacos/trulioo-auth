package dao

import (
	"errors"
	"log"

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
	d.db.Model()
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
