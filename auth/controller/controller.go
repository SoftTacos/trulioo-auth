package controller

import (
	"context"
	"errors"
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt"
	d "github.com/softtacos/trulioo-auth/auth/dao"
	v1 "github.com/softtacos/trulioo-auth/grpc/users/v1"
)

const (
	defaultJwtDuration = time.Hour * 12
)

var (
	defaultHMACSecret = []byte("")
)

func NewAuthController(dao d.Dao, usersClient v1.UsersServiceClient) AuthController {
	return &authController{
		dao:         dao,
		usersClient: usersClient,

		jwtDuration: defaultJwtDuration,
	}
}

type AuthController interface {
	CreateAccount(ctx context.Context, email, password string) (jwt string, err error)
	Login(ctx context.Context, uuid, password string) (jwt string, err error)
}

type authController struct {
	dao         d.Dao
	usersClient v1.UsersServiceClient

	jwtDuration time.Duration
}

func (c *authController) Login(ctx context.Context, uuid, password string) (jwt string, err error) {
	// validate the uuid and password
	if err = c.validateLogin(uuid, password); err != nil {
		log.Println("invalid request: ", err.Error())
		return
	}

	// check if the user exists
	var getUserResponse *v1.GetUserResponse
	getUserResponse, err = c.usersClient.GetUser(ctx, &v1.GetUserRequest{
		Uuid: uuid,
	})
	if err != nil {
		log.Println("failed to retrieve user: ", err.Error())
		return
	}

	jwt, err = c.generateToken(ctx, getUserResponse.GetUser())
	if err != nil {
		return
	}

	// refresh?

	return
}

func (c *authController) generateToken(ctx context.Context, user *v1.User) (tokenString string, err error) {
	// generate a new JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":      user.Uuid,
		"expiresAt": time.Now().Add(c.jwtDuration).Unix(),
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

func (c *authController) CreateAccount(ctx context.Context, email, password string) (jwt string, err error) {
	// validate the uuid and password
	if err = c.validatePassword(password); err != nil {
		log.Println("invalid request: ", err.Error())
		return
	}
	var createUserResponse *v1.CreateUserResponse
	createUserResponse, err = c.usersClient.CreateUser(ctx, &v1.CreateUserRequest{
		Email: email,
	})
	if err != nil {
		log.Println("failed to create user: ", err.Error())
		return
	}

	jwt, err = c.generateToken(ctx, createUserResponse.GetUser())
	if err != nil {
		return
	}

	// refresh?

	return
}

func (c *authController) validatePassword(password string) (err error) {
	if password == "" {
		return errors.New("no password provided")
	}
	return
}
