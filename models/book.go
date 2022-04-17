package models

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

// CreateBookInput includes mandatory keys needed to create a book
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// UpdateBookInput for updating record
type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
