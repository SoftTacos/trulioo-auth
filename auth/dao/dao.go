package dao

import (
	"errors"
	"log"

	gopg "github.com/go-pg/pg/v10"
)

func NewDao(db *gopg.DB) Dao {
	return &dao{
		db: db,
	}
}

type Dao interface {
	GetPassword(uuid string) (passwordHash string, err error)
}

type dao struct {
	db *gopg.DB
}

func (d *dao) GetPassword(uuid string) (passwordHash string, err error) {
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
