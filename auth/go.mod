module github.com/softtacos/trulioo-auth/auth

go 1.17

replace github.com/softtacos/trulioo-auth/grpc => ../grpc

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/softtacos/trulioo-auth/grpc v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/newrelic/go-agent/v3 v3.12.0 // indirect
	github.com/newrelic/go-agent/v3/integrations/nrgrpc v1.3.1 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.45.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
