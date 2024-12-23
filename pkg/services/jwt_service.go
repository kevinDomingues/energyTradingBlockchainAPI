package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
	secretKey string
	issure    string
}

func NewJWTService() *jwtService {
	return &jwtService{
		secretKey: "secret-ket", //adicionar .env
		issure:    "user-api",
	}
}

type Claim struct {
	Sum             string `json:"sum"`
	BlockchainToken string `json:"btkn"`
	UserType        int    `json:"ut"`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(id string, blockchainToken string, userType int) (string, error) {
	claim := &Claim{
		id,
		blockchainToken,
		userType,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid token: %v", token)
		}

		return []byte(s.secretKey), nil
	})

	return parsedToken, err
}
