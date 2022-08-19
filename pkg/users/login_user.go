package users

import (
	"strconv"
	"time"

	"github.com/berrybab6/MovieGo/pkg/common/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password []byte `json:"-"`
}

func (h handler) LoginUser(c *fiber.Ctx) error {
	body := LoginRequestBody{}

	err := c.BodyParser(&body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var user models.User

	const SecretKey = "ThisIsMySecretKey"

	// if result := h.DB.Where("email = ?", body.Email).First(&user); result.Error != nil {
	// 	return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	// }
	h.DB.Where("email = ?", body.Email).First(&user)
	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	if err2 := bcrypt.CompareHashAndPassword(user.Password, body.Password); err2 != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	})

	token, err3 := claims.SignedString([]byte(SecretKey))

	if err3 != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login Succesfully",
		"token":   &token,
	})
}
