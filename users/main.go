package main

import(
	"context"
	c "github.com/softtacos/trulioo-auth/users/controller"
)

func main() {
	var (
		controller c.UsersController
	)
	controller = c.NewUsersController(nil)
	controller.GetUsers(context.Background(),[]uint64{1234})
}
