package controller

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mockDao "github.com/softtacos/trulioo-auth/auth/dao/mocks"
	v1 "github.com/softtacos/trulioo-auth/grpc/users/v1"
	v1mock "github.com/softtacos/trulioo-auth/grpc/users/v1/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_authController_createAccount(t *testing.T) {
	var(
		jwtDuration =time.Hour
		hmacSecret  =[]byte("big ole secret")
	)
	type args struct {
		ctx      context.Context
		email    string
		password string
	}
	tests := []struct {
		name     string
		args     args
		wantUser *v1.User
		wantErr  error
	}{
		{
			name:"happy path",
	args:args{
		ctx:context.Background(),
	}			
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			dao := mockDao.NewMockAuthDao(ctrl)
			usersClient := v1mock.NewMockUsersServiceClient(ctrl)

			c := &authController{
				dao:         dao,
				usersClient: usersClient,
				jwtDuration:jwtDuration,
				hmacSecret: hmacSecret,
			}

			gotUser, err := c.createAccount(tt.args.ctx, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("authController.createAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantUser, gotUser, "user does not match")
		})
	}
}
