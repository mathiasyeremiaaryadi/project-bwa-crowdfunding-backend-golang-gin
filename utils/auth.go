package utils

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

var SECRET_KEY = []byte("s3cr3t-k3y")

type JWT interface {
	GenerateToken(int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewJwtService() JWT {
	return &jwtService{}
}

func (services *jwtService) GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["USER_ID"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signiedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signiedToken, err
	}

	return signiedToken, nil
}

func (services *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	validatedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	return validatedToken, nil
}
