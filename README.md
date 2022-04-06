# Overview
- There are two services, auth and users. Proto files are defined in the grpc directory. Util has helpful infrastructure code that is used by both services
- I opted not to implement a refresh token since that isn't much more complex than the jwt access token, but I was low on time. Adding one would be fairly easy and require extra calls to create a token and storing the token in the DB
# Setup
- Install postgres
- Install golang
- Install goose: https://github.com/pressly/goose

- Each service has a Makefile, to setup the database, cd into the service's directory and run `make setup`
