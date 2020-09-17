package server

import (
	"context"
	"github.com/leshachaplin/grpc-server-library/internal/service"
	"github.com/leshachaplin/grpc-server-library/internal/types"
	"github.com/leshachaplin/grpc-server-library/protocol"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Rpc service.BookService
}

func (s *Server) GetAllBooks(ctx context.Context, req *protocol.EmptyRequest) (*protocol.GetAllBooksResponse, error) {
	books, err := s.Rpc.GetAllBooks(ctx)
	if err != nil {
		log.Errorf("books cannot find: %s", err)
		return nil, err
	}

	var Books = make([]*protocol.Book, 0)

	for i := 0; i < len(books); i++ {
		book := &protocol.Book{
			Name:   books[i].Name,
			Author: books[i].Author,
			Genre:  books[i].Genre,
			Year:   books[i].Year,
		}
		Books = append(Books, book)
	}

	responce := &protocol.GetAllBooksResponse{
		Books: Books,
	}

	return responce, nil
}

func (s *Server) GetBookByName(ctx context.Context, req *protocol.GetBookByNameRequest) (*protocol.GetBookResponse, error) {
	book, err := s.Rpc.GetBookByName(ctx, req.Name)
	if err != nil {
		log.Errorf("book not found by name: %s", err)
		return nil, err
	}
	response := &protocol.GetBookResponse{Book: &protocol.Book{
		Name:   book.Name,
		Author: book.Author,
		Genre:  book.Genre,
		Year:   book.Year,
	}}
	return response, nil
}

func (s *Server) GetBookByAuthor(ctx context.Context, req *protocol.GetBooksByAuthorRequest) (*protocol.GetBookResponse, error) {
	book, err := s.Rpc.GetBookByAuthor(ctx, req.Author)
	if err != nil {
		log.Errorf("book not found by author: %s", err)
		return nil, err
	}
	response := &protocol.GetBookResponse{Book: &protocol.Book{
		Name:   book.Name,
		Author: book.Author,
		Genre:  book.Genre,
		Year:   book.Year,
	}}
	return response, nil
}

func (s *Server) DeleteBook(ctx context.Context, req *protocol.DeleteBookRequest) (*protocol.EmptyResponse, error) {
	err := s.Rpc.DeleteBook(ctx, req.Name)
	if err != nil {
		log.Errorf("book not deleted: %s", err)
		return nil, err
	}
	return &protocol.EmptyResponse{}, nil
}

func (s *Server) AddBook(ctx context.Context, req *protocol.AddBookRequest) (*protocol.EmptyResponse, error) {
	err := s.Rpc.AddBook(ctx, types.Book{
		Name:   req.Book.Name,
		Author: req.Book.Author,
		Genre:  req.Book.Genre,
		Year:   req.Book.Year,
	})
	if err != nil {
		log.Errorf("book not added: %s", err)
		return nil, err
	}
	return &protocol.EmptyResponse{}, nil
}