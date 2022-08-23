package models

type Author struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Biography string `json:"biography"`
	Book      []Book `json:"books"`
}
