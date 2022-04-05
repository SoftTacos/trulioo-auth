package grpc

//go:generate buf mod update
//go:generate buf generate --template buf.gen.yaml
//go:generate buf generate --template buf.gen.tag.yaml
//go:generate mockgen -source=auth/v1/auth_grpc.pb.go  -destination=auth/v1/mocks/mock_Auth.go -package=mocks
//go:generate mockgen -source=users/v1/users_grpc.pb.go  -destination=users/v1/mocks/mock_Users.go -package=mocks
