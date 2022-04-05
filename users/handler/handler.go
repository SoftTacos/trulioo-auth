package handler

import (
	"context"
	"log"

	v1 "github.com/softtacos/trulioo-auth/grpc/users/v1"
	c "github.com/softtacos/trulioo-auth/users/controller"
)

func NewUsersHandler(controller c.UsersController) *UsersHandler {
	return &UsersHandler{
		controller: controller,
	}
}

type UsersHandler struct {
	v1.UnimplementedUsersServiceServer
	controller c.UsersController
}

func (h *UsersHandler) GetUser(ctx context.Context, request *v1.GetUserRequest) (response *v1.GetUserResponse, err error) {

	user, err := h.controller.GetUser(ctx, request.GetEmail())
	if err != nil {
		log.Println("failed to retrieve user:", err.Error())
		return
	}
	response = &v1.GetUserResponse{
		User: user.User,
	}
	return
}

func (h *UsersHandler) CreateUser(ctx context.Context, request *v1.CreateUserRequest) (response *v1.CreateUserResponse, err error) {
	user, err := h.controller.CreateUser(ctx, request.GetEmail())
	if err != nil {
		log.Println("failed to create user:", err.Error())
		return
	}
	response = &v1.CreateUserResponse{
		User: user.User,
	}
	return
}
