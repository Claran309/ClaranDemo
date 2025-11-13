package jwt_util

import (
	"github.com/golang-jwt/jwt/v5"
)

type Util interface {
	GenerateToken(userID int, username string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	ExtractClaims(token *jwt.Token) (jwt.MapClaims, error)
}
