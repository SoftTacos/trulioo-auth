package controller

import (
	"strings"
	"testing"

	v1 "github.com/softtacos/trulioo-auth/grpc/users/v1"
	m "github.com/softtacos/trulioo-auth/users/model"
	"github.com/stretchr/testify/assert"
)

func Test_usersController_validateNewUser(t *testing.T) {
	tests := []struct {
		name    string
		user    m.User
		wantErr error
	}{
		{
			name: "good email",
			user: m.User{
				User: &v1.User{
					Email: "email@email.email",
				},
			},
			wantErr: nil,
		},
		{
			name: "no email",
			user: m.User{
				User: &v1.User{
					Email: "",
				},
			},
			wantErr: errNoEmail,
		},
		{
			name: "email too long",
			user: m.User{
				User: &v1.User{
					Email: strings.Repeat("a", 320) + "@email.com",
				},
			},
			wantErr: errEmailTooLong,
		},
		{
			name: "not an email",
			user: m.User{
				User: &v1.User{
					Email: "this is not an email",
				},
			},
			wantErr: errInvalidAddress,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &usersController{}
			err := c.validateNewUser(tt.user)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
