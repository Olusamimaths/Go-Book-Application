package repository

import "github.com/olusamimaths/go-book-application/src/domain"

type DBHandler interface {
	FindAllBooks() ([]*domain.Book, error)
	SaveBook(book domain.Book) error
	SaveAuthor(author domain.Author) error
}