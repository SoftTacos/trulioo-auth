package controller

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errNoEmail        = status.Error(codes.InvalidArgument, "no email provided")
	errEmailTooLong   = status.Error(codes.InvalidArgument, "email is longer than maximum email length of 320")
	errInvalidAddress = status.Error(codes.InvalidArgument, "invalid email address")
)
