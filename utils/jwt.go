package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT[K any](data map[string]K, secretKey string, expirationTime time.Duration) (string, error) {
	claims := jwt.MapClaims{}

	for key, value := range data {
		claims[key] = value
	}

	claims["exp"] = time.Now().Add(expirationTime).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}

func VerifyJWT[K any](s string, secretKey string) (map[string]K, error) {
	token, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		// validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// return the secret key for validation
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	// check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	// extract the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}
	// convert the claims to the desired type
	data := make(map[string]K)
	for key, value := range claims {
		if key == "exp" {
			continue
		}
		data[key] = value.(K)
	}
	return data, nil
}
