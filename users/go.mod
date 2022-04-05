module github.com/softtacos/trulioo-auth/users

go 1.17

replace github.com/softtacos/trulioo-auth/grpc => ../grpc

replace github.com/softtacos/trulioo-auth/users => ../users

require (
	github.com/go-pg/pg v8.0.7+incompatible
	github.com/softtacos/trulioo-auth/grpc v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.19.0 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.45.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	mellium.im/sasl v0.2.1 // indirect
)
