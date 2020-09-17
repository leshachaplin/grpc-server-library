package service

import (
	"context"
	"github.com/leshachaplin/grpc-server-library/internal/repository"
	"github.com/leshachaplin/grpc-server-library/internal/types"
)

type BookService struct {
	books repository.Books
}

func New(bookServise repository.BookRepository) *BookService {
	return &BookService{
		books: &bookServise,
	}
}

func (s *BookService) GetAllBooks(ctx context.Context) ([]types.Book, error) {
	books, err := s.books.GetAllBooks(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *BookService) GetBookByName(ctx context.Context, name string) (*types.Book, error) {
	book, err := s.books.GetBookByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *BookService) GetBookByAuthor(ctx context.Context, author string) (*types.Book, error) {
	book, err := s.books.GetBookByAuthor(ctx, author)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *BookService) DeleteBook(ctx context.Context, name string) error {
	err := s.books.DeleteBook(ctx, name)
	if err != nil {
		return err
	}
	return nil
}

func (s *BookService) AddBook(ctx context.Context, book types.Book) error {
	err := s.books.AddBook(ctx, book)
	if err != nil {
		return err
	}
	return nil
}
