package models

import (
	"github.com/ayushwar/go/5_bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// ✅ Capitalized struct name so it can be used outside the package
type Book struct {
	gorm.Model
	Name    string `json:"name"`
	Author  string `json:"author"`
	Publish string `json:"publish"`
}

// ✅ init function to connect DB and auto-migrate the Book table
func init() {
	config.Connect() // ✅ Corrected function name (capital 'C')
	db = config.GetDB() // ✅ Corrected function name (capital 'G')
	db.AutoMigrate(&Book{})
}

// ✅ Exported function name: CreateBook instead of createbook
func (b *Book) CreateBook() *Book {
	db.Create(&b) // ✅ Removed deprecated NewRecord
	return b
}

// ✅ Exported function name: GetAllBooks instead of getallbook
func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

// ✅ Exported + renamed: GetBookByID
func GetBookByID(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", id).Find(&getBook)
	return &getBook, db
}

// ✅ Fixed delete logic and exported: DeleteBook
func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID = ?", id).Delete(&book) // ✅ Fixed condition and passed pointer
	return book
}
