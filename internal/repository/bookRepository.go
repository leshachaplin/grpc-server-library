package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/leshachaplin/grpc-server-library/internal/types"
)

type BookRepository struct {
	db *sqlx.DB
}

func NewBookRepository(database sqlx.DB) *BookRepository {
	return &BookRepository{
		db: &database,
	}
}

func (r *BookRepository) GetBookByName(ctx context.Context, name string) (*types.Book, error) {
	rows, err := r.db.QueryxContext(ctx, `SELECT name, author, genre, year FROM "book" WHERE name = $1`, name)
	if err != nil {
		return nil, err
	}
	book := types.Book{}
	for rows.Next() {
		err := rows.StructScan(&book)
		_ = err
	}
	return &book, err
}

func (r *BookRepository) GetBookByAuthor(ctx context.Context, author string) (*types.Book, error) {
	rows, err := r.db.QueryxContext(ctx, `SELECT name, author, genre, year FROM "book" WHERE author = $1`, author)
	if err != nil {
		return nil, err
	}
	book := types.Book{}
	for rows.Next() {
		err := rows.StructScan(&book)
		_ = err
	}
	return &book, err
}

func (r *BookRepository) GetAllBooks(ctx context.Context) ([]types.Book, error) {
	rows, err := r.db.QueryxContext(ctx, `SELECT * FROM "book"`)
	if err != nil {
		return nil, err
	}
	book := types.Book{}
	var books = make([]types.Book, 0)
	for rows.Next() {
		err := rows.StructScan(&book)
		books = append(books, book)
		_ = err
	}
	return books, err
}

func (r *BookRepository) DeleteBook(ctx context.Context, name string) error {
	_, err := r.db.QueryContext(ctx, `delete from "book" where name = $1`, name)
	return err
}

func (r *BookRepository) AddBook(ctx context.Context, book types.Book) error {
	_, err := r.db.QueryContext(ctx, `INSERT into "book" (name, author, genre, year) values ($1, $2, $3, $4)`,
		book.Name, book.Author, book.Genre, book.Year)
	return err
}
