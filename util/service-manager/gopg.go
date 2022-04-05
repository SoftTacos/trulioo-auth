package util

import (
	"time"

	gopg "github.com/go-pg/pg/v10"
)

func CreateGoPgDB(url string) (*gopg.DB, error) {
	options, err := gopg.ParseURL(url)
	if err != nil {
		return nil, err
	}
	options.DialTimeout = 20 * time.Second
	db := gopg.Connect(options)

	// check connection
	_, err = db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	return db, nil
}
