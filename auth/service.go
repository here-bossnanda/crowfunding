package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)


type Service interface {
	GenerateToken(userId int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
} 

var SECRET_KEY = []byte("CROWFUNDING_s3cr3t_k3Y")

func (s *jwtService) GenerateToken(userId int) (string, error) {
	payload := jwt.MapClaims{}
	payload["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	decodedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error){
		_,ok := token.Method.(*jwt.SigningMethodHMAC) 
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return decodedToken, err
	}

	return decodedToken, nil
}