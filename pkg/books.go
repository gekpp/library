package booksapi

import (
	"context"
	"database/sql"
	books "library/gen/books"
	"log"
)

// books service example implementation.
// The example methods log the requests and return zero values.
type bookssrvc struct {
	logger *log.Logger
	db     *sql.DB
}

// NewBooks returns the books service implementation.
func NewBooks(logger *log.Logger) books.Service {
	return &bookssrvc{logger, nil}
}

// List books
func (s *bookssrvc) List(ctx context.Context) (res *books.LibraryBooks, err error) {
	rows, err := s.db.Query("SELECT id, title, annotation, author, images FROM Books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []*books.Book
	for rows.Next() {
		var b books.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Annotation, &b.Author, &b.Images); err != nil {
			return nil, err
		}
		data = append(data, &b)
	}
	res = &books.LibraryBooks{Data: data}
	s.logger.Print("books.list")
	return res, nil
}

// Mark book as reserved. Once a book is reserved timer starts with timeout for
// the book to become picked up. Timeout is configurable.
// Once timeout is expired book becomes available
func (s *bookssrvc) Reserve(ctx context.Context, p *books.ReservePayload) (res interface{}, err error) {
	// 1. get caller username
	username := ctx.Value(contextKeyUsername).(string)

	s.logger.Printf("books.reserve: User %q reserves book(id=%q)", username, p.BookID)

	return "ok", nil
}

// Mark book as picked up
func (s *bookssrvc) Pickup(ctx context.Context, p *books.PickupPayload) (res interface{}, err error) {
	s.logger.Print("books.pickup")
	return
}

// Mark book as returned
func (s *bookssrvc) Return(ctx context.Context, p *books.ReturnPayload) (res interface{}, err error) {
	s.logger.Print("books.return")
	return
}

// Subscribe the caller on the next 'book's become available
func (s *bookssrvc) Subscribe(ctx context.Context, p *books.SubscribePayload) (res interface{}, err error) {
	s.logger.Print("books.subscribe")
	return
}
