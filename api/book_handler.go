package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"my-book-app/internal/book"
	"my-book-app/storage"
)

type BookHandler struct {
	repository storage.PostgresBookRepository
}

func NewBookHandler(repository storage.PostgresBookRepository) *BookHandler {
	return &BookHandler{
			repository: repository,
	}
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
    var newBook book.Book
    err := json.NewDecoder(r.Body).Decode(&newBook)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    err = h.repository.Create(&newBook)
    if err != nil {
        http.Error(w, "Failed to create book", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newBook)
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
			http.Error(w, "Invalid book ID", http.StatusBadRequest)
			return
	}

	updatedBook := &book.Book{ID: bookID}
	err = json.NewDecoder(r.Body).Decode(updatedBook)
	if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
	}

	err = h.repository.Update(updatedBook)
	if err != nil {
			http.Error(w, "Failed to update book", http.StatusInternalServerError)
			return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedBook)
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
			http.Error(w, "Invalid book ID", http.StatusBadRequest)
			return
	}

	err = h.repository.Delete(bookID)
	if err != nil {
			http.Error(w, "Failed to delete book", http.StatusInternalServerError)
			return
	}

	w.WriteHeader(http.StatusNoContent)
}


func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	bookID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
			http.Error(w, "Invalid book ID", http.StatusBadRequest)
			return
	}

	book, err := h.repository.Read(bookID)
	if err != nil {
			http.Error(w, "Failed to fetch book", http.StatusInternalServerError)
			return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
