package dao

import (
	"context"
	"log"

	gopg "github.com/go-pg/pg/v10"
	m "github.com/softtacos/trulioo-auth/users/model"
	errutil "github.com/softtacos/trulioo-auth/util/error"
)

type UsersDao interface {
	GetUser(ctx context.Context, email string) (user m.User, err error)
	CreateUser(ctx context.Context, user m.User) (m.User, error)
}

func NewUsersDao(db *gopg.DB) UsersDao {
	return &usersDao{
		db:     db,
		errMap: dberrmap,
	}
}

type usersDao struct {
	db     *gopg.DB
	errMap errutil.DbErrorMap
}

func (d *usersDao) GetUser(ctx context.Context, email string) (user m.User, err error) {
	err = d.db.Model(&user).Where("email = ?", email).Select()
	if err != nil {
		log.Println("failed to retrieve user: ", err)
		if err == gopg.ErrNoRows {
			err = errutil.ErrDoesNotExist("user", "email")
		} else {
			err = errutil.HandlePgErr(d.errMap, err)
		}
	}
	return
}

func (d *usersDao) CreateUser(ctx context.Context, user m.User) (m.User, error) {
	_, err := d.db.Model(&user).Insert()
	if err != nil {
		log.Println("failed to create user: ", err)
		err = errutil.HandlePgErr(d.errMap, err)
	}
	return user, err
}
