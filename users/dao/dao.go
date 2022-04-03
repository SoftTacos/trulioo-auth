package dao

import(
	"context"

	m "github.com/softtacos/trulioo-auth/users/model"
)
type UsersDao interface {
	GetUsers(ctx context.Context, ids []uint64)(users []m.User,err error)
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

func (d *usersDao)GetUsers(ctx context.Context,ids []uint64)(users []m.User,err error){

	return
}

func (d *usersDao)CreateUser(ctx context.Context,user m.User)(m.User, error){

	return user,nil
}