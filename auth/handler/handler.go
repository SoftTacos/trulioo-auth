package handler

import (
	"context"
	"log"

	c "github.com/softtacos/trulioo-auth/auth/controller"
	v1 "github.com/softtacos/trulioo-auth/grpc/auth/v1"
)

func NewAuthHandler(controller c.AuthController) *AuthHandler {
	return &AuthHandler{
		controller: controller,
	}
}

type AuthHandler struct {
	v1.UnimplementedAuthServiceServer
	controller c.AuthController
}

func (h *AuthHandler) Login(ctx context.Context, request *v1.LoginRequest) (response *v1.LoginResponse, err error) {
	jawt, err := h.controller.Login(ctx, request.Email, request.Password)
	if err != nil {
		log.Println(err)
		return
	}
	response = &v1.LoginResponse{
		Jwt: jawt,
	}
	return
}

func (h *AuthHandler) Signup(ctx context.Context, request *v1.SignupRequest) (response *v1.SignupResponse, err error) {
	jawt, err := h.controller.CreateAccount(ctx, request.Email, request.Password)
	if err != nil {
		log.Println(err)
		return
	}
	response = &v1.SignupResponse{
		Jwt: jawt,
	}
	return
}
