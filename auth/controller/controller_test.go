package controller

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	mockCtrl "github.com/softtacos/trulioo-auth/auth/controller/mocks"
	mockDao "github.com/softtacos/trulioo-auth/auth/dao/mocks"
	v1 "github.com/softtacos/trulioo-auth/grpc/users/v1"
	v1mock "github.com/softtacos/trulioo-auth/grpc/users/v1/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func Test_authController_createAccount(t *testing.T) {
	var (
		testError   = errors.New("failed to do a thing")
		jwtDuration = time.Hour
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

func Test_authController_generateToken(t *testing.T) {
	tests := []struct {
		name           string
		user           *v1.User
		generateResult string
		generateErr    error
		wantErr        error
	}{
		{
			name: "happy",
			user: &v1.User{
				Uuid:  "definitely a uuid!",
				Email: "def an email",
			},
			generateResult: "big ole JWT token string",
			generateErr:    nil,
			wantErr:        nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			tokenGenerator := mockCtrl.NewMockTokenGenerator(ctrl)
			c := &authController{
				jwtDuration:    time.Hour,
				tokenGenerator: tokenGenerator,
			}
			matcher := gomock.Matcher(ClaimsMatcher{
				timeTolerance: time.Second,
				claims: jwt.MapClaims{
					"uuid":      tt.user.Uuid,
					"email":     tt.user.Email,
					"expiresAt": time.Now().Add(c.jwtDuration).Unix(),
				},
			})
			tokenGenerator.EXPECT().Generate(jwt.SigningMethodHS256, matcher).Return(tt.generateResult, tt.generateErr)

			gotTokenString, err := c.generateToken(tt.user)

			if tt.wantErr != nil {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.generateResult, gotTokenString)
		})
	}
}

type ClaimsMatcher struct {
	timeTolerance time.Duration
	claims        jwt.MapClaims // this is just a map[string]interface{}
}

func (m ClaimsMatcher) String() string {
	return fmt.Sprintf("matches claims: %v", m.claims)
}

func (m ClaimsMatcher) Matches(arg interface{}) (matches bool) {
	actualClaims, ok := arg.(jwt.MapClaims)
	if !ok {
		return
	}

	// get each key and value in the actual claims passed in
	for actualKey, actualValue := range actualClaims {
		// check if the keys supplied exist in the expected claims map
		if expectedValue, exists := m.claims[actualKey]; !exists {
			return
		} else { // if both maps do have the same keys
			actualTime, ok := actualValue.(time.Time)
			if ok {
				expectedTime := expectedValue.(time.Time)
				diff := actualTime.Sub(expectedTime)
				if diff < 0 {
					diff = diff * -1
				}
				if diff > m.timeTolerance {
					return
				}
				continue
			}
			if !reflect.DeepEqual(expectedValue, actualValue) {
				return
			}
		}
	}

	return true
}
