package http

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	bookRepository "github.com/kenethrrizzo/bookland-service/cmd/api/data/books"
	bookDomain "github.com/kenethrrizzo/bookland-service/cmd/api/domain/books"
	bookHandler "github.com/kenethrrizzo/bookland-service/cmd/api/router/http/books"
)

const (
	BOOK_BASE_URL = "/books"
)

func NewHTTPHandler(db *sql.DB) http.Handler {
	router := chi.NewRouter()

	/* Book handlers */
	bookRepo := bookRepository.NewStore(db)
	bookService := bookDomain.NewService(bookRepo)
	bookHandler := bookHandler.NewHandler(bookService)

	router.Get(fmt.Sprintf("%s/get-books", BOOK_BASE_URL), bookHandler.GetAllBooks)
	router.Get(fmt.Sprintf("%s/get-books/{bookID}", BOOK_BASE_URL), bookHandler.GetBookByID)
	router.Post(fmt.Sprintf("%s/register-book", BOOK_BASE_URL), bookHandler.RegisterNewBook)

	return router
}