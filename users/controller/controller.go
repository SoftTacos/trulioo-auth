package controller

import (
	"context"
	"log"
	"regexp"
	"time"

	v1 "github.com/softtacos/trulioo-auth/grpc/users/v1"

	//v1 "github.com/softtacos/trulioo-auth/grpc/users"
	"github.com/google/uuid"
	d "github.com/softtacos/trulioo-auth/users/dao"
	m "github.com/softtacos/trulioo-auth/users/model"
)

const maxEmailLength = 320

var (
	// this is simplified the real email regex is a terrifying monstrosity: https://stackoverflow.com/questions/201323/how-can-i-validate-an-email-address-using-a-regular-expression
	emailRegex = regexp.MustCompile(`^[a-z0-9]*@[a-z0-9]*\.[a-z]*$`)
)

type UsersController interface {
	GetUser(ctx context.Context, email string) (user m.User, err error)
	CreateUser(ctx context.Context, email string) (user m.User, err error)
}

func NewUsersController(dao d.UsersDao) UsersController {
	return &usersController{
		dao: dao,
	}
}

type usersController struct {
	dao d.UsersDao
}

func (c *usersController) GetUser(ctx context.Context, email string) (user m.User, err error) {
	user, err = c.dao.GetUser(ctx, email)
	if err != nil {
		log.Println("failed to retrieve user:", err)
	}
	return
}

func (c *usersController) CreateUser(ctx context.Context, email string) (user m.User, err error) {
	user = c.generateUser(email)

	if err = c.validateNewUser(user); err != nil {
		log.Println("invalid user: ", err)
		return
	}

	_, err = c.dao.CreateUser(ctx, user)
	if err != nil {
		log.Println("failed to create user: ", err)
		return
	}
	return
}

func (c *usersController) generateUser(email string) (user m.User) {
	user.User = &v1.User{
		Uuid:  uuid.New().String(),
		Email: email,
	}
	now := time.Now().UTC()
	user.CreatedAt = &now
	return
}

func (c *usersController) validateNewUser(user m.User) (err error) {
	if user.Email == "" {
		err = errNoEmail
	} else if len(user.Email) > maxEmailLength {
		err = errEmailTooLong
	} else if !emailRegex.MatchString(user.Email) {
		err = errInvalidAddress
	}
	return
}
