package controller

import (
	"context"
	"errors"
	"testing"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	mockCtrl "github.com/softtacos/trulioo-auth/auth/controller/mocks"
	mockDao "github.com/softtacos/trulioo-auth/auth/dao/mocks"
	v1 "github.com/softtacos/trulioo-auth/grpc/users/v1"
	v1mock "github.com/softtacos/trulioo-auth/grpc/users/v1/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_authController_createAccount(t *testing.T) {
	var (
		testError   = errors.New("failed to do a thing")
		jwtDuration = time.Hour
		hmacSecret  = []byte("big ole secret")
	)
	type args struct {
		ctx      context.Context
		email    string
		password string
	}
	tests := []struct {
		name                string
		args                args
		createUserResponse  *v1.CreateUserResponse
		createUserError     error
		hashResult          []byte
		hashError           error
		createPasswordError error
		wantErr             error
	}{
		{
			name: "happy path",
			args: args{
				ctx:      context.Background(),
				email:    "email@email.email",
				password: "very secret password",
			},
			createUserResponse: &v1.CreateUserResponse{
				User: &v1.User{
					Email: "email@email.email",
					Uuid:  uuid.New().String(),
				},
			},
			createUserError:     nil,
			hashResult:          []byte("the hash slinging slasher"),
			hashError:           nil,
			createPasswordError: nil,
			wantErr:             nil,
		},
		{
			name: "create user failure",
			args: args{
				ctx:      context.Background(),
				email:    "email@email.email",
				password: "very secret password",
			},
			createUserResponse:  nil,
			createUserError:     testError,
			hashResult:          nil,
			hashError:           testError,
			createPasswordError: nil,
			wantErr:             testError,
		},
		{
			name: "create user failure",
			args: args{
				ctx:      context.Background(),
				email:    "email@email.email",
				password: "very secret password",
			},
			createUserResponse: &v1.CreateUserResponse{
				User: &v1.User{
					Email: "email@email.email",
					Uuid:  uuid.New().String(),
				},
			},
			createUserError:     nil,
			hashResult:          nil,
			hashError:           nil,
			createPasswordError: testError,
			wantErr:             testError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			dao := mockDao.NewMockAuthDao(ctrl)
			usersClient := v1mock.NewMockUsersServiceClient(ctrl)
			pwHasher := mockCtrl.NewMockPasswordHasher(ctrl)
			c := &authController{
				dao:         dao,
				usersClient: usersClient,
				pwHasher:    pwHasher,
				jwtDuration: jwtDuration,
				hmacSecret:  hmacSecret,
			}

			usersClient.EXPECT().CreateUser(tt.args.ctx, &v1.CreateUserRequest{
				Email: tt.args.email,
			}).Return(tt.createUserResponse, tt.createUserError)

			if tt.createUserError == nil {
				pwHasher.EXPECT().HashAndSalt([]byte(tt.args.password), bcrypt.MinCost).Return(tt.hashResult, tt.hashError)
			}
			if tt.hashError == nil {
				dao.EXPECT().CreatePassword(tt.createUserResponse.User.Uuid, string(tt.hashResult)).Return(tt.createPasswordError)
			}
			gotUser, err := c.createAccount(tt.args.ctx, tt.args.email, tt.args.password)
			if err != nil {
				require.NotNil(t, tt.wantErr)
				require.Equal(t, tt.wantErr.Error(), err.Error())
				return
			}
			assert.Equal(t, tt.createUserResponse.User, gotUser, "user does not match")
		})
	}
}
