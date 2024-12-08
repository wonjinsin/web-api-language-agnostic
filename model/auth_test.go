package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestSignup_Validate(t *testing.T) {
	name := "testuser"
	tests := []struct {
		name     string
		signup   Signup
		expected bool
	}{
		{
			name: "valid signup",
			signup: Signup{
				Email:    "test@example.com",
				Password: "password123",
				Name:     &name,
			},
			expected: true,
		},
		{
			name: "empty email",
			signup: Signup{
				Email:    "",
				Password: "password123",
				Name:     &name,
			},
			expected: false,
		},
		{
			name: "empty password",
			signup: Signup{
				Email:    "test@example.com",
				Password: "",
				Name:     &name,
			},
			expected: false,
		},
		{
			name: "nil nick",
			signup: Signup{
				Email:    "test@example.com",
				Password: "password123",
				Name:     nil,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.signup.Validate())
		})
	}
}

func TestSignin_Validate(t *testing.T) {
	tests := []struct {
		name     string
		signin   Signin
		expected bool
	}{
		{
			name: "valid signin",
			signin: Signin{
				Email:    "test@example.com",
				Password: "password123",
			},
			expected: true,
		},
		{
			name: "empty email",
			signin: Signin{
				Email:    "",
				Password: "password123",
			},
			expected: false,
		},
		{
			name: "empty password",
			signin: Signin{
				Email:    "test@example.com",
				Password: "",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.signin.Validate())
		})
	}
}

func TestSignin_CheckPassword(t *testing.T) {
	var password Password = "password123"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	assert.NoError(t, err)

	tests := []struct {
		name     string
		signin   Signin
		hash     string
		expected bool
	}{
		{
			name: "correct password",
			signin: Signin{
				Email:    "test@example.com",
				Password: password,
			},
			hash:     string(hash),
			expected: true,
		},
		{
			name: "incorrect password",
			signin: Signin{
				Email:    "test@example.com",
				Password: "wrongpassword",
			},
			hash:     string(hash),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.signin.CheckPassword(tt.hash))
		})
	}
}

func TestSignin_String(t *testing.T) {
	tests := []struct {
		name     string
		signin   Signin
		expected string
	}{
		{
			name: "normal password",
			signin: Signin{
				Email:    "test@example.com",
				Password: "password123",
			},
			expected: "Email[test@example.com] Password[p*********3]",
		},
		{
			name: "short password",
			signin: Signin{
				Email:    "test@example.com",
				Password: "pw",
			},
			expected: "Email[test@example.com] Password[pw]",
		},
		{
			name: "three character password",
			signin: Signin{
				Email:    "test@example.com",
				Password: "pwd",
			},
			expected: "Email[test@example.com] Password[p*d]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.signin.String())
		})
	}
}
