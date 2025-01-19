package models

import (
	"github.com/rohitbisht/go-book-management-system/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()

	if db != nil {
		db.AutoMigrate(&Book{})
	}
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books

}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(&book)
	return book
}

func UpdateBook(book *Book) (*Book, error) {
	if err := db.Save(book).Error; err != nil {
		return nil, err
	}

	return book, nil
}
