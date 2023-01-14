package usecases

import (
	"log"

	"github.com/olusamimaths/go-book-application/src/domain"
)

type BookInteractor struct {
	BookRepository domain.BookRepository
}

func NewInteractor(repository domain.BookRepository) BookInteractor {
	return BookInteractor{repository}
}

func (interactor *BookInteractor) CreateBook(book domain.Book) error {
	err := interactor.BookRepository.SaveBook(book)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (interactor *BookInteractor) FindAll() ([]*domain.Book, error) {
	books, err := interactor.BookRepository.FindAll()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return books, nil
}