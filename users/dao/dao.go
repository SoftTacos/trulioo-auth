package dao

import (
	"context"
	"log"

	gopg "github.com/go-pg/pg"
	m "github.com/softtacos/trulioo-auth/users/model"
)

type UsersDao interface {
	GetUser(ctx context.Context, email string) (user m.User, err error)
	CreateUser(ctx context.Context, user m.User) (m.User, error)
}

func NewUsersDao(db *gopg.DB) UsersDao {
	return &usersDao{
		db: db,
	}
}

type usersDao struct {
	db *gopg.DB
}

func (d *usersDao) GetUser(ctx context.Context, email string) (user m.User, err error) {
	err = d.db.Model(&user).Where("email = ?", email).Select()
	if err != nil {
		log.Println("failed to retrieve user: ", err)
	}
	return
}

func (d *usersDao) CreateUser(ctx context.Context, user m.User) (m.User, error) {
	_, err := d.db.Model(&user).Insert()
	if err != nil {
		log.Println("failed to create user: ", err)
	}
	return user, err
}
