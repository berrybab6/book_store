package books

import (
	"github.com/berrybab6/MovieGo/pkg/users"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// handler MovieGo
type handler struct {
	DB *gorm.DB
}

// RegisterRoutes godoc
func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := app.Group("books")
	routes.Post("/", h.AddBook)
	routes.Get("/all", users.ValidateJWT, h.GetBooks)

	routes.Get("/getbook/:id", h.GetBook)
	routes.Put("/updatebooksbyid/:id", h.UpdateBook)
	routes.Delete("/deletebook/:id", h.DeleteBook)
	routes.Put("/upload/:id", h.Upload)
}
