package main

import (
	"context"
	"os"

	c "github.com/softtacos/trulioo-auth/auth/controller"
	clients "github.com/softtacos/trulioo-auth/grpc"
	v1 "github.com/softtacos/trulioo-auth/grpc/users"
)

const (
	usersUrlEnv = "USERS_CLIENT_ADDRESS"
	dbUrlEnv    = "DB_URL"
)

func main() {
	os.Setenv("USERS_CLIENT_ADDRESS", ":11000")

	clientManager := clients.NewClientManager()

	authClientConn := clientManager.Create(os.Getenv(usersUrlEnv))
	usersClient := v1.NewUsersServiceClient(authClientConn)

	controller := c.NewAuthController(usersClient)
	controller.GenerateToken(context.Background(), "", "")

	// TODO: add shutdown on interrupt
}
