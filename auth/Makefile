.PHONY: all

all: setup migrate

setup:
	dropdb tl_auth || true
	createdb tl_auth --owner=postgres
	cd migrations && goose postgres "user=postgres dbname=tl_auth sslmode=disable" up

create:
	cd migrations && goose create $(name) sql

upgrade-local:
	goose -dir migrations postgres "user=postgres dbname=tl_auth sslmode=disable" up

downgrade-local:
	goose -dir migrations postgres "user=postgres dbname=tl_auth sslmode=disable" down
