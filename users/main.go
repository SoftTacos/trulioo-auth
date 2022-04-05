package main

import (
	"os"

	v1 "github.com/softtacos/trulioo-auth/grpc/users/v1"
	c "github.com/softtacos/trulioo-auth/users/controller"
	d "github.com/softtacos/trulioo-auth/users/dao"
	h "github.com/softtacos/trulioo-auth/users/handler"
	sm "github.com/softtacos/trulioo-auth/util/service-manager"
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
	var (
		dao        d.UsersDao
		controller c.UsersController
		handler    *h.UsersHandler
	)

	karen := &sm.ServiceManager{
		ServiceSetup: func(m *sm.ServiceManager) (err error) {
			controller = c.NewUsersController(dao)
			handler = h.NewUsersHandler(controller)
			v1.RegisterUsersServiceServer(m.CreateGrpcServer(os.Getenv(grpcPortEnv)), handler)
			return
		},
		DatabaseSetup: func(m *sm.ServiceManager) (err error) {
			db := m.CreateDbConnection(os.Getenv(dbUrlEnv))
			dao = d.NewUsersDao(db)
			return
		},
	}

	defer karen.Stop()
	karen.Start()

}
