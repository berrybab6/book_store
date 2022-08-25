package models

// Book MovieGo
type Book struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Title    string `json:"title"`
	AuthorID uint   `json:"authorid"`
	// Author      Author `gorm:"foreignKey:AuthorID;association_foreignkey:ID;;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Description string `json:"description"`
	Attachment  string `json:"attachment"`
}
