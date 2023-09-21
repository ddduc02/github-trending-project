package security

import (
	"time"

	"github.com/ddduc02/gh-trending/models"
	"github.com/dgrijalva/jwt-go"
)

const SECRET_KEY = "ducdodeptrai"

func GenerateToken(user models.User) (string, error) {
	claims := models.JwtCustomClaims{
		UserId: user.UserId,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
