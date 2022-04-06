package dao

import (
	gopg "github.com/go-pg/pg/v10"
	errutil "github.com/softtacos/trulioo-auth/util/error"
)

var dberrmap = errutil.DbErrorMap{
	"users": map[string]func(gopg.Error) error{
		"23505": func(pgErr gopg.Error) error {
			return errutil.ErrAlreadyExists("user")
		},
	},
}
