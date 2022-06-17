package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string)(*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte("MOMEN_money_management")

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	// token header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signToken, err
	}

	return signToken, nil
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error)  {
	tokenParse, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return tokenParse, err
	}
	return tokenParse, nil
}