package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/olusamimaths/go-book-application/src/domain"
	"github.com/olusamimaths/go-book-application/src/usecases"
)

type AuthorController struct {
	authorInteractor usecases.AuthorInteractor
}

func NewAuthorController(authorInteractor usecases.AuthorInteractor) *AuthorController {
	return &AuthorController{authorInteractor}
}

func (controller *AuthorController) Add(res http.ResponseWriter, req http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var author domain.Author
	err := json.NewDecoder(req.Body).Decode(&author)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}

	err = controller.authorInteractor.CreateAuthor(author)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
}