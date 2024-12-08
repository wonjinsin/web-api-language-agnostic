package model

import (
	"os"
	"path/filepath"
	"pikachu/util"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func getTestKeys(t *testing.T) ([]byte, []byte) {
	rootDir := util.GetRootDir() // reuse the existing function from user.go

	prvKey, err := os.ReadFile(filepath.Join(rootDir, "mock-local", "token_key"))
	assert.NoError(t, err, "Failed to read private key")

	pubKey, err := os.ReadFile(filepath.Join(rootDir, "mock-local", "token_key.pub"))
	assert.NoError(t, err, "Failed to read public key")

	return prvKey, pubKey
}

func TestNewToken(t *testing.T) {
	tests := []struct {
		name        string
		accessToken string
		expected    *Token
	}{
		{
			name:        "valid token",
			accessToken: "valid-token",
			expected: &Token{
				AccessToken: "valid-token",
			},
		},
		{
			name:        "empty token",
			accessToken: "",
			expected: &Token{
				AccessToken: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewToken(tt.accessToken)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNewUserClaim(t *testing.T) {
	projectName := "test-project"
	domain := "test.com"
	email := "test@example.com"
	uid := "test-uid"
	user := &User{
		UID:   uid,
		Email: email,
	}

	claim := NewUserClaim(projectName, domain, user)

	assert.NotNil(t, claim)
	assert.NotEmpty(t, claim.ID)
	assert.Equal(t, projectName, claim.Issuer)
	assert.Equal(t, uid, claim.Subject)
	assert.Equal(t, jwt.ClaimStrings{util.TokenAudienceAccount}, claim.Audience)
	assert.NotNil(t, claim.ExpiresAt)
	assert.NotNil(t, claim.IssuedAt)

	expectedExpiry := claim.IssuedAt.Add(24 * time.Hour)
	assert.WithinDuration(t, expectedExpiry, claim.ExpiresAt.Time, time.Second)

	assert.Equal(t, util.TokenTypeBearer, claim.Type)
	assert.Equal(t, projectName, claim.AuthorizedParty)
	assert.Equal(t, email, claim.Email)
	assert.Equal(t, []string{"User"}, claim.Roles)
	assert.Equal(t, []string{"Email"}, claim.Scope)
}

func TestTokenClaim_GenerateToken(t *testing.T) {
	prvKey, pubKey := getTestKeys(t)

	tests := []struct {
		name      string
		claim     *TokenClaim
		wantError bool
	}{
		{
			name: "valid token generation",
			claim: &TokenClaim{
				RegisteredClaims: jwt.RegisteredClaims{
					ID:        "test-id",
					Issuer:    "test-project",
					Subject:   "test-user",
					Audience:  []string{util.TokenAudienceAccount},
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
					IssuedAt:  jwt.NewNumericDate(time.Now()),
				},
				Type:            util.TokenTypeBearer,
				AuthorizedParty: "test-project",
				Email:           "test@example.com",
				Roles:           []string{"User"},
				Scope:           []string{"Email"},
			},
			wantError: false,
		},
		{
			name: "invalid private key",
			claim: &TokenClaim{
				Type: util.TokenTypeBearer,
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tokenKey []byte
			if tt.name == "invalid private key" {
				tokenKey = []byte("invalid-key")
			} else {
				tokenKey = prvKey
			}

			token, err := tt.claim.GenerateToken(tokenKey)

			if tt.wantError {
				assert.Error(t, err)
				assert.Empty(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)

				parsedPubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKey)
				assert.NoError(t, err)

				parsedToken, err := jwt.ParseWithClaims(token, &TokenClaim{}, func(token *jwt.Token) (interface{}, error) {
					return parsedPubKey, nil
				})
				assert.NoError(t, err)
				assert.True(t, parsedToken.Valid)
			}
		})
	}
}

func TestTokenClaim_GetUserUID(t *testing.T) {
	tests := []struct {
		name     string
		claim    *TokenClaim
		expected string
	}{
		{
			name: "valid user UID",
			claim: &TokenClaim{
				RegisteredClaims: jwt.RegisteredClaims{
					Subject: "test-uid",
				},
			},
			expected: "test-uid",
		},
		{
			name: "empty user UID",
			claim: &TokenClaim{
				RegisteredClaims: jwt.RegisteredClaims{
					Subject: "",
				},
			},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.claim.GetUserUID()
			assert.Equal(t, tt.expected, result)
		})
	}
}
