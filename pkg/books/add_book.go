package books

import (
	"github.com/berrybab6/MovieGo/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type AddBookRequestBody struct {
	Title       string `gorm:"title"`
	AuthorID    uint   `gorm:"authorid"`
	Description string `gorm:"description"`
}

func (h handler) AddBook(c *fiber.Ctx) error {
	body := AddBookRequestBody{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var book models.Book

	book.Title = body.Title
	book.AuthorID = body.AuthorID
	book.Description = body.Description

	if result := h.DB.Create(&book); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())

	}
	return c.Status(fiber.StatusCreated).JSON(&book)
}
