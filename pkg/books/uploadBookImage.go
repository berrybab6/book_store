package books

import (
	"fmt"

	"github.com/berrybab6/MovieGo/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

// Upload godoc
// @Summary      Upload Book attachment
// @Description  Upload file
// @Tags         books
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path      uint   true  "Book ID"
// @Param        file  formData  file  true  "book attachement"
// @Success      200   {object}  controller.Message
// @Failure      400   {object}  httputil.HTTPError
// @Failure      404   {object}  httputil.HTTPError
// @Failure      500   {object}  httputil.HTTPError
// @Router       /books/upload/{id} [put]
func (h handler) Upload(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// if err != nil {
	// 	return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able to process your expense"}})
	// }

	file, err := ctx.FormFile("attachment")

	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able upload your attachment"}})
	}

	ctx.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))

	var book models.Book
	h.DB.First(&book, id)
	h.DB.Model(&book).Update("attachment", file.Filename)

	return ctx.JSON(fiber.Map{"message": "Attachment uploaded successfully"})
}
