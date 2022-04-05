package controller

import (
	"context"
	"log"


	auth "github.com/softtacos/trulioo-auth/grpc/auth"
	//v1 "github.com/softtacos/trulioo-auth/grpc/users"
	d "github.com/softtacos/trulioo-auth/users/dao"
)

type UsersController interface{
	GetUsers(ctx context.Context,ids []uint64)(err error)
	CreateUser(ctx context.Context,email,password string)(err error)
}

func NewUsersController(dao d.UsersDao,authClient auth.AuthServiceClient) UsersController{
	return &usersController{
		dao:		dao,
		authClient: authClient,
	}
}

type usersController struct {
	dao d.UsersDao
	authClient auth.AuthServiceClient
}

func (c *usersController)GetUsers(ctx context.Context,ids []uint64)(err error){
	log.Println("GET USERS")
	return
}

func (c *usersController)CreateUser(ctx context.Context,email,password string)(err error) {
	log.Println("GET USERS")

	// check if user with email already exists

	// login(generate JWT)


	return
}
