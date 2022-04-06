package controller

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errLoginFailure = status.Error(codes.InvalidArgument, "invalid email and password combination")
)
