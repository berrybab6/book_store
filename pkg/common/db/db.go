package db

import (
	"log"

	"github.com/berrybab6/MovieGo/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Init MovieGo
func Init(url string) *gorm.DB {

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
