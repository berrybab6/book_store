package users

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var jwtSecret = []byte("movieGoSuperSecretKey")

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

type Authorization struct {
	Authorization string `reqHeader:"Authorization"`
}

func GenerateJWT(email string, username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(time.Hour * 1)

	claims := &JWTClaim{
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtSecret)
	return
}

func ValidateJWT(c *fiber.Ctx) error {
	bearer := new(Authorization)

	if err := c.ReqHeaderParser(bearer); err != nil {
		return fiber.ErrBadRequest
	}
	var tokenVal string

	if val := strings.Split(bearer.Authorization, " "); len(val) < 2 {
		fmt.Println(len(val))
		return fiber.ErrForbidden
	} else {
		tokenVal = val[1]
	}

	token, err := jwt.ParseWithClaims(
		tokenVal,
		&JWTClaim{},
		func(*jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

	// fmt.Println(token)
	claims, ok := token.Claims.(*JWTClaim)
	if !(ok && token.Valid) {
		return fiber.ErrForbidden
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("Token Expired")
		return err
	}
	if err != nil {
		fmt.Println(err)
		return fiber.ErrForbidden
	}

	return c.Next()

}
