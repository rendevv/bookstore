package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rendevv/bookstore/backend/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"column:name" json:"name"`
	Author      string `gorm:"column:author" json:"author"`
	Publication string `gorm:"column:publication" json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var book Book
	result := db.Where("id = ?", id).First(&book)
	return &book, result
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("id = ?", id).Delete(&book)
	return book
}
