package users

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterUserRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := app.Group("users")
	routes.Post("/", h.CreateUser)
	// routes.Get("/", h.GetUsers)
	routes.Post("/login", h.LoginUser)

	secured := app.Group("/secured").Use(ValidateJWT)
	{
		secured.Get("/", h.GetUsers)
		secured.Put("/:id", h.ChangePassword)
		secured.Put("/forgot", h.ForgotPassword)

		// secured.GET("/ping", controllers.Ping)
	}
}
