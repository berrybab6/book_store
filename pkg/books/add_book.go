package books

import (
	"github.com/berrybab6/MovieGo/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

// AddBookRequestBody MovieGo
type AddBookRequestBody struct {
	Title       string `gorm:"title"`
	AuthorID    uint   `gorm:"authorid"`
	Description string `gorm:"description"`
}

// AddBook godoc
// @Summary      Add a Book
// @Description  add by json book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book  body      AddBookRequestBody  true  "Add Book"
// @Success      201      {object}  models.Book
// @Failure      400      {object}  httputil.HTTPError
// @Failure      404      {object}  httputil.HTTPError
// @Failure      500      {object}  httputil.HTTPError
// @Router       /books [post]
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
