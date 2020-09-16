package repository

import (
	"context"
	"github.com/leshachaplin/grpc-server-library/internal/types"
)

type Books interface {
	GetAllBooks(ctx context.Context) ([]types.Book, error)
	GetBookByName(ctx context.Context, name string) (*types.Book, error)
	GetBookByAuthor(ctx context.Context, author string) (*types.Book, error)
	AddBook(ctx context.Context, book types.Book) error
	DeleteBook(ctx context.Context, name string) error
}
