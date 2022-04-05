package controller

import (
	"reflect"
	"testing"

	d "github.com/softtacos/trulioo-auth/users/dao"
	m "github.com/softtacos/trulioo-auth/users/model"
)

func Test_usersController_generateUser(t *testing.T) {
	type fields struct {
		dao d.UsersDao
	}
	type args struct {
		email string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantUser m.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &usersController{
				dao: tt.fields.dao,
			}
			if gotUser := c.generateUser(tt.args.email); !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("usersController.generateUser() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}
