module github.com/softtacos/trulioo-auth/auth

go 1.17

replace github.com/softtacos/trulioo-auth/grpc => ../grpc

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/softtacos/trulioo-auth/grpc v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.45.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)