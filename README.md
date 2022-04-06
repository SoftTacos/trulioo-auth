# Overview
There are two services, auth and users. Proto files are defined in the grpc directory. Util has helpful infrastructure code that is used by both services

# Setup
- Install postgres
- Install golang
- Install goose: https://github.com/pressly/goose

- Each service has a Makefile, to setup the database, cd into the service's directory and run `make setup`
