package authors

import (
	"github.com/berrybab6/MovieGo/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type AddAuthorRequestBody struct {
	FirstName string `gorm:"first_name"`
	LastName  string `gorm:"last_name"`
	Email     string `gorm:"email"`
	Biography string `gorm:"biography"`
}

func (h handler) AddAuthor(c *fiber.Ctx) error {
	body := AddAuthorRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var author models.Author

	author.FirstName = body.FirstName
	author.LastName = body.LastName
	author.Email = body.Email
	author.Biography = body.Biography

	if result := h.DB.Create(&author); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "A new Author Registered Succesfully",
		"author":  &author,
	})
}
