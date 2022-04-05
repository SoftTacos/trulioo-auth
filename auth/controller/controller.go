package controller

import (
	"context"
	"errors"
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt"
	v1 "github.com/softtacos/trulioo-auth/grpc/users"
)

const (
	defaultJwtDuration = time.Hour * 12
)

var (
	defaultHMACSecret = []byte("")
)

func NewAuthController() AuthController {
	return &authController{
		usersClient: nil,
		jwtDuration: defaultJwtDuration,
	}
}

type AuthController interface {
	GenerateToken(ctx context.Context, email, password string) (jwt string, err error)
}

type authController struct {
	usersClient v1.UsersServiceClient

	jwtDuration time.Duration
}

func (c *authController) GenerateToken(ctx context.Context, uuid, password string) (jwt string, err error) {
	// validate the uuid and password
	if err = c.validateLogin(uuid, password); err != nil {
		log.Println("invalid request: ", err.Error())
		return
	}

	jwt, err = c.generateToken(ctx, uuid, password)
	if err != nil {
		return
	}

	// refresh?

	return
}

func (c *authController) generateToken(ctx context.Context, uuid, password string) (tokenString string, err error) {
	// check if the user exists
	var getUserResponse *v1.GetUserResponse
	getUserResponse, err = c.usersClient.GetUser(ctx, &v1.GetUserRequest{
		Uuid: uuid,
	})
	if err != nil {
		log.Println("failed to retrieve user: ", err.Error())
		return
	}

	// generate a new JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid": getUserResponse.User.Uuid,
		"nbf":  time.Now().Add(c.jwtDuration).Unix(),
	})

	tokenString, err = token.SignedString(defaultHMACSecret)
	if err != nil {
		log.Println("failed to retrieve user: ", err.Error())
		return
	}
	return
}

func (c *authController) validateLogin(uuid, password string) (err error) {
	if uuid == "" {
		return errors.New("no uuid provided")
	}
	if password == "" {
		return errors.New("no password provided")
	}
	return
}
