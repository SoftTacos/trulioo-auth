module github.com/softtacos/trulioo-auth/users

go 1.17

replace github.com/softtacos/trulioo-auth/grpc => ../grpc

replace github.com/softtacos/trulioo-auth/users => ../users

require (
	github.com/go-pg/pg v8.0.7+incompatible
	github.com/google/uuid v1.1.2
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.19.0 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405 // indirect
	mellium.im/sasl v0.2.1 // indirect
)
