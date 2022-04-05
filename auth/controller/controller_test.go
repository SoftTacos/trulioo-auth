package controller

import (
	"context"
	"reflect"
	"testing"
	"time"

	d "github.com/softtacos/trulioo-auth/auth/dao"
	v1 "github.com/softtacos/trulioo-auth/grpc/users/v1"
)

func Test_authController_CreateAccount(t *testing.T) {
	type fields struct {
		dao         d.AuthDao
		usersClient v1.UsersServiceClient
		jwtDuration time.Duration
		hmacSecret  []byte
	}
	type args struct {
		ctx      context.Context
		email    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantJwt string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &authController{
				dao:         tt.fields.dao,
				usersClient: tt.fields.usersClient,
				jwtDuration: tt.fields.jwtDuration,
				hmacSecret:  tt.fields.hmacSecret,
			}
			gotJwt, err := c.CreateAccount(tt.args.ctx, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("authController.CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotJwt != tt.wantJwt {
				t.Errorf("authController.CreateAccount() = %v, want %v", gotJwt, tt.wantJwt)
			}
		})
	}
}

func Test_authController_createAccount(t *testing.T) {
	type fields struct {
		dao         d.AuthDao
		usersClient v1.UsersServiceClient
		jwtDuration time.Duration
		hmacSecret  []byte
	}
	type args struct {
		ctx      context.Context
		email    string
		password string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantUser *v1.User
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &authController{
				dao:         tt.fields.dao,
				usersClient: tt.fields.usersClient,
				jwtDuration: tt.fields.jwtDuration,
				hmacSecret:  tt.fields.hmacSecret,
			}
			gotUser, err := c.createAccount(tt.args.ctx, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("authController.createAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("authController.createAccount() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}
