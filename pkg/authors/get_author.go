package authors

import (
	"github.com/berrybab6/MovieGo/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetAuthor(c *fiber.Ctx) error {
	id := c.Params("id")
	var author models.Author
	if result := h.DB.Preload("Book").First(&author, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	return c.Status(fiber.StatusOK).JSON(&author)
}
