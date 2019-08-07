package booksapi

import (
	"fmt"
	books "library/gen/books"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/google/uuid"
)

const (
	statusFree     = "free"
	statusReserved = "reserved"
	statusPickedUp = "pickedup"
	statusReturned = "returned"
)

// BookStatus defines actual book status, who changed it and when
type BookStatus struct {
	BookID int64      `db:"book_id"`
	Status string     `db:"status"`
	Who    *uuid.UUID `db:"who"`
	When   time.Time  `db:"when"`
}

// BooksRepo defines Books repository interface
type BooksRepo interface {
	GetAllBooks() ([]*books.Book, error)
	GetBookStatus(...int64) (map[int64]BookStatus, error)
	ChangeBookStatus(bookID int64, subscriberID uuid.UUID, status string, createdBy uuid.UUID) error
}

// dbBooksRepo implements BooksRepo with database backend
type dbBooksRepo struct {
	db *sqlx.DB
}

// NewDbBooksRepo create new repository with database backend
func NewDbBooksRepo(db *sqlx.DB) BooksRepo {
	repo := dbBooksRepo{db}
	return &repo
}

func (r *dbBooksRepo) GetAllBooks() ([]*books.Book, error) {
	query := `SELECT id, title, annotation, author, images FROM book`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*books.Book, 0)
	for rows.Next() {
		var (
			b      books.Book
			images stringArray
		)
		if err := rows.Scan(&b.ID, &b.Title, &b.Annotation, &b.Author, &images); err != nil {
			return nil, err
		}
		b.Images = images
		res = append(res, &b)
	}
	return res, nil
}

func (r *dbBooksRepo) GetBookStatus(bookIDs ...int64) (map[int64]BookStatus, error) {
	if len(bookIDs) == 0 {
		return make(map[int64]BookStatus), nil
	}

	query := `SELECT DISTINCT ON (book_id) book_id, status, subscriber_id as who, created_at as when
	FROM book_status_change_event 
	WHERE 
		book_id IN (?)
	ORDER BY book_id ASC, created_at DESC`

	var ids []int64
	ids = bookIDs
	query, args, err := sqlx.In(query, ids)

	if err != nil {
		return nil, fmt.Errorf("could not construct IN query: %v", err)
	}
	query = r.db.Rebind(query)
	fmt.Println(query)

	var statuses []BookStatus
	if err := r.db.Select(&statuses, query, args...); err != nil {
		return nil, fmt.Errorf("could not select IN query: %v", err)
	}

	res := make(map[int64]BookStatus)
	for _, id := range bookIDs {
		status, ok := res[id]
		if !ok {
			res[id] = BookStatus{BookID: id, Status: statusFree}
		} else {
			res[id] = status
		}
	}

	return res, nil
}

func (r *dbBooksRepo) ChangeBookStatus(bookID int64, subscriberID uuid.UUID, status string, createdBy uuid.UUID) error {
	query := `INSERT INTO book_status_change_event (book_id, subscriber_id, status, created_by)
	VALUES
	($1, $2, $3, $4)`
	_, err := r.db.Exec(query, bookID, subscriberID, status, createdBy)
	return err
}
