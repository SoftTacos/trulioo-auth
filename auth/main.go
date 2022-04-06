package main

import (
	"os"

	c "github.com/softtacos/trulioo-auth/auth/controller"
	d "github.com/softtacos/trulioo-auth/auth/dao"
	h "github.com/softtacos/trulioo-auth/auth/handler"
	v1 "github.com/softtacos/trulioo-auth/grpc/auth/v1"
	uv1 "github.com/softtacos/trulioo-auth/grpc/users/v1"
	sm "github.com/softtacos/trulioo-auth/util/service-manager"
)

const (
	usersUrlEnv  = "USERS_CLIENT_ADDRESS"
	dbUrlEnv     = "DB_URL"
	grpcPortEnv  = "GRPC_PORT"
	jwtSecretEnv = "JWT_SECRET"
)

// using this to fake the envs that would normally be set in a chart
// init gets called before anything else
func init() {
	os.Setenv(dbUrlEnv, "postgres://postgres:postgres@localhost/tl_auth?sslmode=disable")
	os.Setenv(usersUrlEnv, ":11001")
	os.Setenv(grpcPortEnv, "11000")
	os.Setenv(jwtSecretEnv, "shhhh it's a secret!")
}

func main() {
	var (
		usersClient uv1.UsersServiceClient
		dao         d.AuthDao
		controller  c.AuthController
		handler     *h.AuthHandler
	)

	karen := &sm.ServiceManager{
		ClientsSetup: func(m *sm.ServiceManager) {
			usersClientConn := m.CreateClientConnection(os.Getenv(usersUrlEnv))
			usersClient = uv1.NewUsersServiceClient(usersClientConn)
		},
		ServiceSetup: func(m *sm.ServiceManager) {
			controller = c.NewAuthController([]byte(os.Getenv(jwtSecretEnv)), dao, usersClient)
			handler = h.NewAuthHandler(controller)
			v1.RegisterAuthServiceServer(m.CreateGrpcServer(os.Getenv(grpcPortEnv)), handler)
		},
		DatabaseSetup: func(m *sm.ServiceManager) {
			db := m.CreateDbConnection(os.Getenv(dbUrlEnv))
			dao = d.NewAuthDao(db)
		},
	}

	defer karen.Stop()
	karen.Start()
}
