package encrypt

import (
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CompareHashAndPassword(hashedPassword, password *string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(*hashedPassword), []byte(*password))
}

func GenerateFromPassword(password *string) (err error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.MinCost)
	if err != nil {
		return
	}
	*password = string(encryptedPassword)
	return
}

func NewTokenWithClaims(claims jwt.Claims) (token string, err error) {
	claimsToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return claimsToken.SignedString([]byte("movie-festival"))
}

func Parse(token string) (tokenRaw string, claims jwt.MapClaims, err error) {
	tokenString := strings.ReplaceAll(token, "Bearer ", "")
	jwtToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte("movie-festival"), nil
	})
	if err != nil {
		return
	}
	return jwtToken.Raw, jwtToken.Claims.(jwt.MapClaims), nil
}
