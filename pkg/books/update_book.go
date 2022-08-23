package books

import (
	"github.com/berrybab6/MovieGo/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type UpdateBookRequestBody struct {
	Title       string `gorm:"title"`
	AuthorID    uint   `gorm:"author_id"`
	Description string `gorm:"description"`
}

func (h handler) UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateBookRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var book models.Book

	result := h.DB.First(&book, id)

	if result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	book.Title = body.Title
	book.Description = body.Description
	book.AuthorID = body.AuthorID

	h.DB.Save(&book)
	return c.Status(fiber.StatusOK).JSON(&book)
}
