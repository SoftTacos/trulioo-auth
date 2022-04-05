package main

import (
	"context"
	"fmt"
	"os"
	"time"

	gopg "github.com/go-pg/pg/v10"
	c "github.com/softtacos/trulioo-auth/auth/controller"
	d "github.com/softtacos/trulioo-auth/auth/dao"
	h "github.com/softtacos/trulioo-auth/auth/handler"
	clients "github.com/softtacos/trulioo-auth/grpc"
	v1 "github.com/softtacos/trulioo-auth/grpc/auth/v1"
	uv1 "github.com/softtacos/trulioo-auth/grpc/users/v1"
)

const (
	usersUrlEnv = "USERS_CLIENT_ADDRESS"
	dbUrlEnv    = "DB_URL"

	defaultMaxPoolSize = 10
)

func main() {
	os.Setenv("USERS_CLIENT_ADDRESS", ":11000")

	var db *gopg.DB
	db = createGoPgDB(dbUrlEnv)
	dao := d.NewDao(db)
	clientManager := clients.NewClientManager()

	authClientConn := clientManager.Create(os.Getenv(usersUrlEnv))
	usersClient := uv1.NewUsersServiceClient(authClientConn)

	controller := c.NewAuthController(dao, usersClient)
	controller.Login(context.Background(), "", "")

	handler := h.NewAuthHandler(controller)
	v1.RegisterAuthServiceServer(nil, handler)
	// TODO: add shutdown on interrupt
}

func createGoPgDB(dbKey string) *gopg.DB {
	url := os.Getenv(dbKey)
	if url == "" {
		panic(fmt.Errorf("%s is not set", dbKey))
	}
	db, err := CreateGoPgDB(url)
	if err != nil {
		panic(err)
	}
	return db
}

func CreateGoPgDB(name string) (*gopg.DB, error) {
	options, err := gopg.ParseURL(name)
	if err != nil {
		return nil, err
	}
	options.DialTimeout = 20 * time.Second
	options.PoolSize = defaultMaxPoolSize

	//check if connection is up
	db := gopg.Connect(options) //this is only structured this way for testing gopg
	_, err = db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	return db, nil
}
