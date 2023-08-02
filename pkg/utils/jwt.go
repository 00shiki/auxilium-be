package utils

import (
	"fmt"
	"github.com/go-chi/jwtauth"
	"time"
)

var TokenAuth *jwtauth.JWTAuth

func InitJWT() {
	TokenAuth = jwtauth.New("HS256", []byte("rahasia"), nil)
	_, tokenString, _ := TokenAuth.Encode(map[string]interface{}{"user_id": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

func CreateToken(userId uint, role int, exp time.Duration) (string, error) {
	now := time.Now().UTC()
	claims := map[string]interface{}{
		"id":   userId,
		"role": role,
		"exp":  now.Add(exp).Unix(),
		"iat":  now.Unix(),
		"nbf":  now.Unix(),
	}
	_, token, err := TokenAuth.Encode(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}
