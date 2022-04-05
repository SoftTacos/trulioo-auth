module github.com/softtacos/trulioo-auth/auth

go 1.17

replace github.com/softtacos/trulioo-auth/grpc => ../grpc

replace github.com/softtacos/trulioo-auth/util => ../util

require (
	github.com/go-pg/pg/v10 v10.10.6
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang/mock v1.6.0
	github.com/softtacos/trulioo-auth/grpc v0.0.0-00010101000000-000000000000
	github.com/softtacos/trulioo-auth/util v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-pg/zerochecker v0.2.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/vmihailenco/bufpool v0.1.11 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.4 // indirect
	github.com/vmihailenco/tagparser v0.1.2 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20210923061019-b8560ed6a9b7 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.45.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
	mellium.im/sasl v0.2.1 // indirect
)
