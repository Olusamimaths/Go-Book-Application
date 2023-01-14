package repository

import "github.com/olusamimaths/go-book-application/src/domain"

type BookRepo struct {
	handler DBHandler
}

func NewBookRepo(handler DBHandler) BookRepo {
	return BookRepo{handler}
}

func (repo BookRepo) SaveBook(book domain.Book) error {
	return repo.handler.SaveBook(book)
}

func (repo BookRepo) FindAll() ([]*domain.Book, error) {
	results, err := repo.handler.FindAllBooks()
	if err != nil {
		return results, err
	}
	return results, nil
}