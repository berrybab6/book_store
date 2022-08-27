package main

import (
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/berrybab6/MovieGo/docs"
	"github.com/berrybab6/MovieGo/pkg/authors"
	"github.com/berrybab6/MovieGo/pkg/books"
	"github.com/berrybab6/MovieGo/pkg/common/config"
	"github.com/berrybab6/MovieGo/pkg/common/db"
	"github.com/berrybab6/MovieGo/pkg/users"
	"github.com/gofiber/fiber/v2"
)

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @securitydefinitions.apikey  JWT
// @in                          header
// @name                        Authorization

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	// tmpl := template.Must(template.ParseGlob(".templates/template.html"))

	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at Config", err)
	}
	app := fiber.New()

	// url := fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=disable",
	// 	c.DBDriver,

	// 	c.DBUser,
	// 	c.DBPassword,
	// 	c.DBHost,
	// 	c.DBPort,
	// 	c.DBName)

	db := db.Init(c.DBUrl)

	// app.Get("/", func(ctx *fiber.Ctx) error {
	// 	return ctx.Status(fiber.StatusOK).SendString(c.Port)
	// })
	books.RegisterRoutes(app, db)
	users.RegisterUserRoutes(app, db)
	authors.RegisterAuthorRoutes(app, db)
	app.Get("/", HealthCheck)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Listen(c.Port)

}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}
