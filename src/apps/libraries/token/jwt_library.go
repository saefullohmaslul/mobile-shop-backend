package token

import (
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims is settings jwt
type Claims struct {
	Payload interface{} `json:"payload"`
	jwt.StandardClaims
}

// SignJWT is method to sign jwt with asymetric
func SignJWT(payload interface{}) (string, time.Time, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Payload: payload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	privateKey, err := ioutil.ReadFile("config/key/private.key")
	if err != nil {
		return "", expirationTime, err
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", expirationTime, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", expirationTime, err
	}

	return tokenString, expirationTime, nil
}
