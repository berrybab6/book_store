package db

import (
	"log"

	"github.com/berrybab6/MovieGo/pkg/common/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Init MovieGo
func Init(url string) *gorm.DB {
	// DB_URL="$}c.DB_DRIVER://DB_USER:DB_PASSWORD@DB_HOST:DB_PORT/DB_NAME"

	// url = fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)

	// lela := c.DBDriver + "://" + c.DBUser + ":" + c.DBPassword + "@" + c.DBHost + ":" + c.DBPort + "/" + c.DBName

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Author{})
	// db.Model(&models.Author{})

	db.Migrator().CreateConstraint(&models.Book{}, "Authors")
	db.Migrator().CreateConstraint(&models.Book{}, "fk_books_author")

	// db.Model(&models.Book{}).AddForeignKey("author_id", "authors(id)", "RESTRICT", "RESTRICT")
	// db.Model(&mode{}).AddForeignKey("profile_id", "profiles(id)", "NO ACTION", "NO ACTION")

	return db
}
