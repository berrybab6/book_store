package users

import (
	"github.com/berrybab6/MovieGo/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h handler) LoginUser(c *fiber.Ctx) error {
	body := LoginRequestBody{}
	// var data map[string]string
	err := c.BodyParser(&body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var user models.User

	// const SecretKey = "ThisIsMySecretKey"

	// if result := h.DB.Where("email = ?", body.Email).First(&user); result.Error != nil {
	// 	return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	// }
	record := h.DB.Where("email = ?", body.Email).First(&user)
	if record.Error != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials Email",
		})
	}

	if err := user.CheckPassword(body.Password); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials Password",
		})
	}

	tokenString, err := GenerateJWT(user.Email, user.Username)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login Succesfully",
		"token":   tokenString,
	})
}
