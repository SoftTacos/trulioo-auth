.PHONY: all

all: setup migrate

setup:
	dropdb tl_users || true
	createdb tl_users --owner=postgres
	cd migrations && goose postgres "user=postgres dbname=tl_users sslmode=disable" up

create:
	cd migrations && goose create $(name) sql

upgrade-local:
	goose -dir migrations postgres "user=postgres dbname=tl_users sslmode=disable" up

downgrade-local:
	goose -dir migrations postgres "user=postgres dbname=tl_users sslmode=disable" down
