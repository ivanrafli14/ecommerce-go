package jwt

import (
	"github.com.ivanrafli14/ecommerce-golang/pkg/config"
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
	"github.com/golang-jwt/jwt/v5"
)

type Interface interface {
	CreateJWTToken(AuthID int, AuthRole string) (string, error)
	VerifyJWTToken(token string) (int, string, error)
}
type jsonWebToken struct {
	secretKey string
}

func NewJWT(cfg config.JWTConfig) Interface {
	secretKey := cfg.SecretKey

	return &jsonWebToken{
		secretKey: secretKey,
	}
}

func (j *jsonWebToken) CreateJWTToken(AuthID int, AuthRole string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"AuthID":   AuthID,
		"AuthRole": AuthRole,
	})
	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *jsonWebToken) VerifyJWTToken(tokenString string) (int, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return 0, "", response.ErrJWTInvalid
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return int(claims["AuthID"].(float64)), claims["AuthRole"].(string), nil
	}
	return 0, "", err
}
