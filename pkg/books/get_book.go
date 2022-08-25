package books

import (
	"github.com/berrybab6/MovieGo/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

// GetBook godoc
// @Summary      Show a Book Detail
// @Description  get string by ID
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id   path      uint  true  "Book ID"
// @Success      200  {object}  models.Book
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /books/getbook/{id} [get]
func (h handler) GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	if result := h.DB.First(&book, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	return c.Status(fiber.StatusOK).JSON(&book)
}
