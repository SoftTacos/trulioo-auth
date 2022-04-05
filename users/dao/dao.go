package dao

import (
	"context"

	gopg "github.com/go-pg/pg"
	v1 "github.com/softtacos/trulioo-auth/grpc/users"
	m "github.com/softtacos/trulioo-auth/users/model"
)

type UsersDao interface {
	GetUsers(ctx context.Context,filters *v1.GetUsersRequest)(users []m.User,err error)
	CreateUser(ctx context.Context, user m.User)(m.User, error)
}

func NewUsersDao(db *gopg.DB)UsersDao{
	return &usersDao{
		db:db,
	}
}

type usersDao struct {
	db *gopg.DB
}

func (d *usersDao)GetUsers(ctx context.Context,filters *v1.GetUsersRequest)(users []m.User,err error){
	if len(filters.Uuids) > 0 {

	}
	return
}

func (d *usersDao)CreateUser(ctx context.Context,user m.User)(m.User, error){

	return user,nil
}