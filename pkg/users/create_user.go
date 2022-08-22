package users

import (
	"github.com/berrybab6/MovieGo/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type CreateUserRequestBody struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h handler) CreateUser(c *fiber.Ctx) error {
	body := CreateUserRequestBody{}

	err := c.BodyParser(&body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// password, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 14)

	var user models.User

	user.Name = body.Name
	user.Username = body.Username
	user.Email = body.Email

	if err := user.HashPassword(body.Password); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())

	}
	h.DB.Create(&user)
	return c.Status(fiber.StatusCreated).JSON(&user)
}
