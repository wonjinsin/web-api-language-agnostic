package middleware

import (
	"context"
	"crypto/rsa"
	"fmt"
	"net/http"
	"pikachu/config"
	"pikachu/model"
	"pikachu/util"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var authNotRequired = map[string]bool{
	"/api/auths/signup": true,
	"/api/auths/signin": true,
}

// AuthMiddleware ...
func AuthMiddleware(conf *config.ViperConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if isPassAuth(c.Request().URL.Path) {
				return next(c)
			}

			tokenString, err := getBearerToken(c)
			if err != nil {
				return response(c, http.StatusUnauthorized, err.Error())
			}

			key, err := getTokenKey(conf)
			if err != nil {
				return response(c, http.StatusUnauthorized, err.Error())
			}

			token, err := jwt.ParseWithClaims(tokenString, &model.TokenClaim{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return key, nil
			})

			if err != nil {
				return response(c, http.StatusUnauthorized, "Invalid token")
			}

			claims, ok := token.Claims.(*model.TokenClaim)
			if !ok || !token.Valid {
				return response(c, http.StatusUnauthorized, "Invalid token claims")
			}

			newReq := c.Request().WithContext(
				context.WithValue(c.Request().Context(), util.UUID, claims.GetUserUID()),
			)
			c.SetRequest(newReq)
			return next(c)
		}
	}
}

func isPassAuth(path string) bool {
	return authNotRequired[path]
}

func getBearerToken(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("token is required")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != util.TokenTypeBearer {
		return "", fmt.Errorf("token is required")
	}

	return parts[1], nil
}

func getTokenKey(conf *config.ViperConfig) (*rsa.PublicKey, error) {
	keyBytes, ok := conf.Get(util.ConfigPubTokenKey).([]byte)
	if !ok {
		return nil, fmt.Errorf("cannot read token key")
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("cannot read token key")
	}

	return key, nil
}
