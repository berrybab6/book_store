package authors

import (
	"github.com/berrybab6/MovieGo/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetAuthors(c *fiber.Ctx) error {
	var authors []models.Author
	result := h.DB.Preload("Book").Find(&authors)

	if result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"authors": &authors,
	})
}
