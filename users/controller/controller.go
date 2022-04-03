package controller

import (
	"context"
	"log"

	auth "github.com/softtacos/trulioo-auth/grpc/auth"
)

type UsersController interface{
	GetUsers(ctx context.Context,ids []uint64)(err error)
	CreateUser(ctx context.Context,email,password string)(err error)
}

func NewUsersController(authClient auth.AuthServiceClient) UsersController{
	return &usersController{
		authClient: authClient,
	}
}

type usersController struct {
	authClient auth.AuthServiceClient
}

func (c *usersController) 	GetUsers(ctx context.Context,ids []uint64)(err error){
	log.Println("GET USERS")
return
}

func (c *usersController)CreateUser(ctx context.Context,email,password string)(err error) {
	return
}
