package repository

import "github.com/olusamimaths/go-book-application/src/domain"

type AuthorRepo struct {
	handler DBHandler
}

func NewAuthorRepo(handler DBHandler) AuthorRepo {
	return AuthorRepo{handler}
} 

func (repo AuthorRepo) SaveAuthor(author domain.Author) error {
	return repo.handler.SaveAuthor(author)
}