package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

func CreateHS256SignedJWT(secret string, payload map[string]interface{}) (string, error) {

	// method of signing the token: this will be populated in header part of token
	token := jwt.New(jwt.SigningMethodHS256)

	// this is the payload
	claims := token.Claims.(jwt.MapClaims)
	for k, v := range payload {
		claims[k] = v
	}

	// final jwt with signed part using secret
	return token.SignedString([]byte(secret))
}

func IsValidHS256JWTSignature(secret string, token string) (bool, error) {
	splitToken := strings.Split(token, ".")
	if len(splitToken) < 3 {
		return false, errors.New("invalid JWT token")
	}
	err := jwt.SigningMethodHS256.Verify(splitToken[0]+"."+splitToken[1], splitToken[2], []byte(secret))

	if err != nil {
		return false, err
	}

	return true, nil
}
