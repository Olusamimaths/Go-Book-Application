package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/olusamimaths/go-book-application/src/domain"
	"github.com/olusamimaths/go-book-application/src/usecases"
)

type BookController struct {
	bookInteractor usecases.BookInteractor
}

func NewBookController(bookInteractor usecases.BookInteractor) *BookController {
	return &BookController{bookInteractor}
}

func (controller *BookController) Add(res http.ResponseWriter, req http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var book domain.Book
	err := json.NewDecoder(req.Body).Decode(&book)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}
	
	err = controller.bookInteractor.CreateBook(book)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (controller *BookController) FindAll(res http.ResponseWriter, req http.Request) {
	res.Header().Set("Content-Type", "application/json")

	results, err := controller.bookInteractor.FindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{err.Error()})
		return
	}

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(results)
}