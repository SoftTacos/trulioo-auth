package controller

//go:generate mockgen -package controller -destination mocks/mock_AuthController.go . AuthController

import (
	"context"
	"errors"
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt"
	d "github.com/softtacos/trulioo-auth/auth/dao"
	v1 "github.com/softtacos/trulioo-auth/grpc/users/v1"
	"golang.org/x/crypto/bcrypt"
)

const (
	defaultJwtDuration = time.Hour
	minPasswordLength  = 8
)

func NewAuthController(secret []byte, dao d.AuthDao, usersClient v1.UsersServiceClient) AuthController {
	return &authController{
		dao:         dao,
		usersClient: usersClient,

		pwHasher: pwHasher{},
		tokenGenerator: &tokenGenerator{
			secret: secret,
		},
		jwtDuration: defaultJwtDuration,
	}
}

type AuthController interface {
	CreateAccount(ctx context.Context, email, password string) (jwt string, err error)
	Login(ctx context.Context, email, password string) (jwt string, err error)
}

type authController struct {
	dao         d.AuthDao
	usersClient v1.UsersServiceClient

	pwHasher       PasswordHasher // mocked the call to hash password for testability
	tokenGenerator TokenGenerator // mocked the call to NewWithClaims for testability
	jwtDuration    time.Duration
}

func (c *authController) Login(ctx context.Context, email, password string) (jwt string, err error) {
	// validate the uuid and password
	if err = c.validateLogin(email, password); err != nil {
		log.Println("invalid request: ", err)
		return
	}

	user, err := c.getUser(ctx, email)
	if err != nil {
		log.Println("failed to retrieve user: ", err)
		return
	}

	jwt, err = c.generateToken(user)
	if err != nil {
		log.Println("failed to geenrate token: ", err)
	}

	return
}

func (c *authController) getUser(ctx context.Context, email string) (user *v1.User, err error) {
	// check if the user exists
	var getUserResponse *v1.GetUserResponse
	getUserResponse, err = c.usersClient.GetUser(ctx, &v1.GetUserRequest{
		Email: email,
	})
	if err != nil {
		log.Println("failed to retrieve user: ", err)
		err = errLoginFailure // don't show the true error, just tell them it failed.
	} else {
		user = getUserResponse.User
	}
	return
}

func (c *authController) CreateAccount(ctx context.Context, email, password string) (jwt string, err error) {
	// validate the uuid and password
	if err = c.validatePassword(password); err != nil {
		log.Println("invalid request: ", err)
		return
	}

	user, err := c.createAccount(ctx, email, password)
	if err != nil {
		log.Println("failed to create an account: ", err)
		return
	}

	jwt, err = c.generateToken(user)
	if err != nil {
		log.Println("failed to generate token: ", err)
	}

	return
}

func (c *authController) createAccount(ctx context.Context, email, password string) (user *v1.User, err error) {
	var createUserResponse *v1.CreateUserResponse
	createUserResponse, err = c.usersClient.CreateUser(ctx, &v1.CreateUserRequest{
		Email: email,
	})
	if err != nil {
		log.Println("failed to create user: ", err)
		return
	}
	user = createUserResponse.User

	// GenerateFromPassword hashes and salts the password
	hash, err := c.pwHasher.HashAndSalt([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return
	}

	err = c.dao.CreatePassword(user.Uuid, string(hash))
	if err != nil {
		log.Println("failed to insert new password: ", err)
	}
	return
}

func (c *authController) generateToken(user *v1.User) (tokenString string, err error) {
	claims := jwt.MapClaims{
		"uuid":      user.GetUuid(),
		"email":     user.GetEmail(),
		"expiresAt": time.Now().Add(c.jwtDuration).Unix(),
	}

	// tokenString, err = c.tokenGenerator.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(c.secret)
	tokenString, err = c.tokenGenerator.Generate(jwt.SigningMethodHS256, claims)
	if err != nil {
		log.Println("failed to retrieve user: ", err)
	}
	return
}

// more robust validation of the email happens on the users service, however we want to make sure the fields were even provided
func (c *authController) validateLogin(email, password string) (err error) {
	if email == "" {
		err = errors.New("no email provided")
	} else if password == "" {
		err = errors.New("no password provided")
	}
	return
}

func (c *authController) validatePassword(password string) (err error) {
	if password == "" {
		err = errors.New("no password provided")
	} else if len(password) < minPasswordLength {
		err = errors.New("password must be at least 8 characters")
	}
	return
}
