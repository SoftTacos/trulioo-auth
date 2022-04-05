package main

import (
	"fmt"
	"log"
	"net"
	"os"

	v1 "github.com/softtacos/trulioo-auth/grpc/users/v1"
	c "github.com/softtacos/trulioo-auth/users/controller"
	d "github.com/softtacos/trulioo-auth/users/dao"
	h "github.com/softtacos/trulioo-auth/users/handler"
	"github.com/softtacos/trulioo-auth/util"
	"google.golang.org/grpc"
)

const (
	dbUrlEnv    = "DB_URL"
	grpcPortEnv = "GRPC_PORT"
)

// using this to fake the envs that would normally be set in a chart
// init gets called before anything else
func init() {
	os.Setenv("GRPC_PORT", "11001")
	os.Setenv("DB_URL", "postgres://postgres:postgres@localhost/tl_users?sslmode=disable")
}

func main() {
	grpcPort := os.Getenv(grpcPortEnv)

	db, err := util.CreateGoPgDB(os.Getenv(dbUrlEnv))
	if err != nil {
		log.Panic("error creating DB connection: ", err.Error())
	}
	dao := d.NewUsersDao(db)

	controller := c.NewUsersController(dao)

	grpcServer := grpc.NewServer()
	grpcListener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	handler := h.NewUsersHandler(controller)
	v1.RegisterUsersServiceServer(grpcServer, handler)
	grpcServer.Serve(grpcListener)

	// TODO: add shutdown on interrupt
}
