package model

import (
	"context"
	"testing"

	"pikachu/util"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestNewUserBySignup(t *testing.T) {
	name := "testuser"
	tests := []struct {
		name    string
		signup  *Signup
		wantErr bool
	}{
		{
			name: "valid signup",
			signup: &Signup{
				Email:    "test@example.com",
				Password: "password123",
				Name:     &name,
			},
			wantErr: false,
		},
		{
			name: "empty signup",
			signup: &Signup{
				Email:    "",
				Password: "",
				Name:     &name,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := NewUserBySignup(tt.signup)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, user)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, user)
			assert.NotEmpty(t, user.UID)
			assert.Equal(t, tt.signup.Email, user.Email)
			assert.NotEmpty(t, user.Password)
			assert.Equal(t, tt.signup.Name, user.Name)
		})
	}
}

func TestUser_AfterFind(t *testing.T) {
	tests := []struct {
		name     string
		user     *User
		isLogin  bool
		wantPass string
	}{
		{
			name: "logged in user",
			user: &User{
				UID:      "test-uid",
				Email:    "test@example.com",
				Password: "hashedpassword",
			},
			isLogin:  true,
			wantPass: "hashedpassword",
		},
		{
			name: "not logged in user",
			user: &User{
				UID:      "test-uid",
				Email:    "test@example.com",
				Password: "hashedpassword",
			},
			isLogin:  false,
			wantPass: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), util.LoginKey, tt.isLogin)
			stmt := &gorm.Statement{Context: ctx}
			db := &gorm.DB{Statement: stmt}

			err := tt.user.AfterFind(db)
			assert.NoError(t, err)
			assert.Equal(t, Password(tt.wantPass), tt.user.Password)
		})
	}
}
