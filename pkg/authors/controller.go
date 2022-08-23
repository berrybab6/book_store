package authors

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterAuthorRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	routes := app.Group("author")
	routes.Post("/", h.AddAuthor)
	routes.Get("/", h.GetAuthors)
	routes.Get("/:id", h.GetAuthor)

}
