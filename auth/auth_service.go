package auth

import (
	"errors"
	"momen/utils"
	"strconv"
	"time"

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



func (s *jwtService) GenerateToken(userID int) (string, error) {
	_,dbConfig := utils.DatabaseSettings()

	SECRET_KEY := []byte(dbConfig.SecretKey)

	tokenHourLifespanString := "24"

	tokenHourLifespan, err := strconv.Atoi(tokenHourLifespanString)

	if err != nil {
		return "", err
	}

	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	claim["exp"] = time.Now().Add(time.Hour * time.Duration(tokenHourLifespan))
	

	// token header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signToken, err
	}

	return signToken, nil
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error)  {
	_,dbConfig := utils.DatabaseSettings()
	SECRET_KEY := []byte(dbConfig.SecretKey)
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