package models

import (
	"github.com/jinzhu/gorm"
	"github.com/zicozydailva/faux-go/bookstore-management-system/pkg/config"
)

var db *gorm.DB

type Book struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db.Where("ID = ?", Id).First(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID = ?", ID).Delete(&book)
	return book
}

func UpdateBook(ID int64, book Book) *Book {
	db.Where("ID = ?", ID).Updates(&book)
	return &book
}

func (b *Book) UpdateBook() *Book {
	db.Save(&b)
	return b
}