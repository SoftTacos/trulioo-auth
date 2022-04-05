package main

import (
	"context"
	c "github.com/softtacos/trulioo-auth/users/controller"
	d "github.com/softtacos/trulioo-auth/users/dao"
)

const(
	dbUrlKey = "DB_URL"
)

func main() {
	var (
		dao        d.UsersDao
		controller c.UsersController
	)
	dao= d.NewUsersDao(nil)


	controller = c.NewUsersController(dao,nil)
	controller.GetUsers(context.Background(),[]uint64{1234})
}
