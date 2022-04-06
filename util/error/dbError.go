package error

import (
	gopg "github.com/go-pg/pg/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DbErrorMap map[string]map[string]func(gopg.Error) error

// pass in a DbErrorMap and the error returned from gopg. The EbErrorMap should be defined on your dao/package just once.
// for a comprehensive list of possible error codes, see: https://www.postgresql.org/docs/13/errcodes-appendix.html
// the most common codes you're going to run into are Integrity Constraint Violation(Class 23) errors
// for a list of the possible fields you can access on the error: https://www.postgresql.org/docs/10/protocol-error-fields.html
func HandlePgErr(errMap DbErrorMap, err error) error {
	if pgErr, ok := err.(gopg.Error); ok {
		tableName := pgErr.Field('t')
		if codeMap, ok := errMap[tableName]; ok {
			errCode := pgErr.Field('C')
			if errFunc, ok := codeMap[errCode]; ok {
				err = errFunc(pgErr)
			}
		}
	}
	return err
}

// common error funcs
func ErrAlreadyExists(typeStr string) error {
	return status.Errorf(codes.AlreadyExists, "a %s with that name already exists", typeStr)
}

func ErrDoesNotExist(typeStr, field string) error {
	return status.Errorf(codes.NotFound, "a %s with that %s does not exist", typeStr, field)
}
