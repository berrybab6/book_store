package books

import (
	"github.com/berrybab6/MovieGo/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	result := h.DB.First(&book, id)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	h.DB.Delete(&book)
	return c.Status(fiber.StatusNoContent).SendString("Deleted Succesfully")
}
