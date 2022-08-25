package books

import (
	"github.com/berrybab6/MovieGo/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

// DeleteBook godoc
// @Summary      Delete a book
// @Description  Delete by book ID
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id   path      uint  true  "Book ID"  Format(int64)
// @Success      204  {object}  model.Book
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.H
// @Router       /books/deletebook/{id} [delete]

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
