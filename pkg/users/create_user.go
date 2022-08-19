package users

import (
	"github.com/berrybab6/MovieGo/pkg/common/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserRequestBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
}

func (h handler) CreateUser(c *fiber.Ctx) error {
	body := CreateUserRequestBody{}

	err := c.BodyParser(&body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	password, _ := bcrypt.GenerateFromPassword(body.Password, 12)

	var user models.User

	user.Name = body.Name
	user.Email = body.Email
	user.Password = password

	h.DB.Create(&user)
	return c.Status(fiber.StatusCreated).JSON(&user)
}
