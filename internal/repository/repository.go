package repository

import (
	"github.com/leshachaplin/grpc-server-library/internal/types"
)

type Books interface {
	GetAllBooks() ([]types.Book, error)
	GetBookByName(name string) (*types.Book, error)
	GetBookByAuthor(author string) (*types.Book, error)
	AddBook(book types.Book) error
	DeleteBook(name string) error
}
