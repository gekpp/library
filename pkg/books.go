package booksapi

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	books "library/gen/books"
	"log"

	goa "goa.design/goa/v3/pkg"

	"github.com/google/uuid"
)

type stringArray []string

func (a stringArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Make the Attrs struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (a *stringArray) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type bookssrvc struct {
	logger *log.Logger
	repo   BooksRepo
}

// NewBooks returns the books service implementation.
func NewBooks(logger *log.Logger, repo BooksRepo) books.Service {
	return &bookssrvc{logger, repo}
}

// List books
func (s *bookssrvc) List(ctx context.Context) (res *books.LibraryBooks, err error) {
	data, err := s.repo.GetAllBooks()
	if err != nil {
		return nil, err
	}

	ids := make([]int64, len(data))
	for i, d := range data {
		ids[i] = d.ID
	}

	statuses, err := s.repo.GetBookStatus(ids...)
	if err != nil {
		s.logger.Printf("couldn't get books statuses: %v", err)
		return nil, goa.PermanentError("internal-error", "Can't get books statuses")
	}
	for _, d := range data {
		d.Status = statuses[d.ID].Status
	}

	res = &books.LibraryBooks{Data: data}
	s.logger.Print("books.list")
	return res, nil
}

// Mark book as reserved. Once a book is reserved timer starts with timeout for
// the book to become picked up. Timeout is configurable.
// Once timeout is expired book becomes available
func (s *bookssrvc) Reserve(ctx context.Context, p *books.ReservePayload) (res interface{}, err error) {
	// get caller userID
	userID := ctx.Value(contextKeyUserID).(uuid.UUID)

	subscriberID, err := uuid.Parse(p.SubscriberID)
	if err != nil {
		return "", goa.PermanentError("bad-request", "can not parse subscriberID value %q as UUID", p.SubscriberID)
	}

	status, err := s.getBookStatus(p.BookID)
	if err != nil {
		return "", goa.PermanentError("internal-error", "can't get book status")
	}

	switch status.Status {
	case statusReserved, statusPickedUp:
		s.logger.Printf("books.reserve: Could not reserve book: book %v is already reserved or pickedup", p.BookID)
		return "", fmt.Errorf("book %q is already reserved or pickedup", p.BookID)
	case statusFree, statusReturned:
		return "", s.repo.ChangeBookStatus(p.BookID, subscriberID, statusReserved, userID)
	default:
		return "", fmt.Errorf("it is impossible to reserve the book because book status is %q", status)
	}
}

// Mark book as picked up
func (s *bookssrvc) Pickup(ctx context.Context, p *books.PickupPayload) (res interface{}, err error) {
	// get caller userID
	userID := ctx.Value(contextKeyUserID).(uuid.UUID)

	payloadSubscriberID, err := uuid.Parse(p.SubscriberID)
	if err != nil {
		return "", goa.PermanentError("bad-request", "can not parse subscriberID value %q as UUID", p.SubscriberID)
	}

	status, err := s.getBookStatus(p.BookID)
	if err != nil {
		return "", goa.PermanentError("internal-error", "can't get book status")
	}

	switch status.Status {
	case statusReserved:
		if *status.Who != payloadSubscriberID {
			return "", fmt.Errorf("Book is reserved by another user. Cancel the reservation first")
		}
		return "", s.repo.ChangeBookStatus(p.BookID, payloadSubscriberID, statusPickedUp, userID)
	case statusFree, statusReturned:
		return "", s.repo.ChangeBookStatus(p.BookID, payloadSubscriberID, statusPickedUp, userID)
	default:
		return "", fmt.Errorf("unsupported book status %q", status)
	}
}

// Mark book as returned
func (s *bookssrvc) Return(ctx context.Context, p *books.ReturnPayload) (res interface{}, err error) {
	// get caller userID
	userID := ctx.Value(contextKeyUserID).(uuid.UUID)

	status, err := s.getBookStatus(p.BookID)
	if err != nil {
		return "", goa.PermanentError("internal-error", "can't get book status")
	}

	newSubscriberID, err := uuid.Parse(p.SubscriberID)
	if err != nil {
		return "", goa.PermanentError("bad-request", "can not parse subscriberID value %q as UUID", p.SubscriberID)
	}

	switch status.Status {
	case statusReserved, statusPickedUp:
		if *status.Who != newSubscriberID {
			return "", fmt.Errorf("Book is reserved or pickedup by another user. Cancel the reservation first")
		}
		return "", s.repo.ChangeBookStatus(p.BookID, newSubscriberID, statusReturned, userID)
	case statusFree, statusReturned:
		return "", goa.PermanentError("bad-request", "Book is not reserved or taken")
	default:
		return "", fmt.Errorf("unsupported book status %q", status)
	}
}

// Subscribe the caller on the next 'book's become available
func (s *bookssrvc) Subscribe(ctx context.Context, p *books.SubscribePayload) (res interface{}, err error) {
	s.logger.Print("books.subscribe")
	return
}

func (s *bookssrvc) getBookStatus(id int64) (BookStatus, error) {
	stss, err := s.repo.GetBookStatus(id)
	if err != nil {
		s.logger.Print(err)
		return BookStatus{}, err
	}

	sts, ok := stss[id]
	if !ok {
		return BookStatus{}, fmt.Errorf("status of the book id=%v not found in repository result", id)
	}

	return sts, nil
}
