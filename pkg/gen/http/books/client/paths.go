// Code generated by goa v3.0.3, DO NOT EDIT.
//
// HTTP request path constructors for the books service.
//
// Command:
// $ goa gen library/design

package client

import (
	"fmt"
)

// ListBooksPath returns the URL path to the books service list HTTP endpoint.
func ListBooksPath() string {
	return "/books/list"
}

// ReserveBooksPath returns the URL path to the books service reserve HTTP endpoint.
func ReserveBooksPath(bookID string) string {
	return fmt.Sprintf("/books/reserve/%v", bookID)
}

// PickupBooksPath returns the URL path to the books service pickup HTTP endpoint.
func PickupBooksPath(bookID string) string {
	return fmt.Sprintf("/books/pickup/%v", bookID)
}

// ReturnBooksPath returns the URL path to the books service return HTTP endpoint.
func ReturnBooksPath(bookID string) string {
	return fmt.Sprintf("/books/return/%v", bookID)
}

// SubscribeBooksPath returns the URL path to the books service subscribe HTTP endpoint.
func SubscribeBooksPath(bookID string) string {
	return fmt.Sprintf("/books/subscribe/%v", bookID)
}
