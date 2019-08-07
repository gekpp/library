// Code generated by goa v3.0.3, DO NOT EDIT.
//
// books HTTP server types
//
// Command:
// $ goa gen library/design

package server

import (
	books "library/gen/books"
	booksviews "library/gen/books/views"

	goa "goa.design/goa/v3/pkg"
)

// ReserveRequestBody is the type of the "books" service "reserve" endpoint
// HTTP request body.
type ReserveRequestBody struct {
	// id of a subscriber picking up the book
	SubscriberID *string `form:"subscriber_id,omitempty" json:"subscriber_id,omitempty" xml:"subscriber_id,omitempty"`
}

// PickupRequestBody is the type of the "books" service "pickup" endpoint HTTP
// request body.
type PickupRequestBody struct {
	// id of a subscriber picking up the book
	SubscriberID *string `form:"subscriber_id,omitempty" json:"subscriber_id,omitempty" xml:"subscriber_id,omitempty"`
}

// ReturnRequestBody is the type of the "books" service "return" endpoint HTTP
// request body.
type ReturnRequestBody struct {
	// id of the Book
	BookID *int64 `form:"book_id,omitempty" json:"book_id,omitempty" xml:"book_id,omitempty"`
	// id of a subscriber returning the book
	SubscriberID *string `form:"subscriber_id,omitempty" json:"subscriber_id,omitempty" xml:"subscriber_id,omitempty"`
}

// ListResponseBody is the type of the "books" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	Data []*BookResponseBody `form:"data" json:"data" xml:"data"`
}

// ReturnInvalidScopesResponseBody is the type of the "books" service "return"
// endpoint HTTP response body for the "invalid-scopes" error.
type ReturnInvalidScopesResponseBody string

// SubscribeInvalidScopesResponseBody is the type of the "books" service
// "subscribe" endpoint HTTP response body for the "invalid-scopes" error.
type SubscribeInvalidScopesResponseBody string

// BookResponseBody is used to define fields on response body types.
type BookResponseBody struct {
	ID         int64  `form:"id" json:"id" xml:"id"`
	Title      string `form:"title" json:"title" xml:"title"`
	Annotation string `form:"annotation" json:"annotation" xml:"annotation"`
	Author     string `form:"author" json:"author" xml:"author"`
	// images are a list of book photos
	Images []string `form:"images" json:"images" xml:"images"`
	Status string   `form:"status" json:"status" xml:"status"`
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "books" service.
func NewListResponseBody(res *booksviews.LibraryBooksView) *ListResponseBody {
	body := &ListResponseBody{}
	if res.Data != nil {
		body.Data = make([]*BookResponseBody, len(res.Data))
		for i, val := range res.Data {
			body.Data[i] = marshalBooksviewsBookViewToBookResponseBody(val)
		}
	}
	return body
}

// NewReturnInvalidScopesResponseBody builds the HTTP response body from the
// result of the "return" endpoint of the "books" service.
func NewReturnInvalidScopesResponseBody(res books.InvalidScopes) ReturnInvalidScopesResponseBody {
	body := ReturnInvalidScopesResponseBody(res)
	return body
}

// NewSubscribeInvalidScopesResponseBody builds the HTTP response body from the
// result of the "subscribe" endpoint of the "books" service.
func NewSubscribeInvalidScopesResponseBody(res books.InvalidScopes) SubscribeInvalidScopesResponseBody {
	body := SubscribeInvalidScopesResponseBody(res)
	return body
}

// NewReservePayload builds a books service reserve endpoint payload.
func NewReservePayload(body *ReserveRequestBody, bookID int64, token string) *books.ReservePayload {
	v := &books.ReservePayload{
		SubscriberID: *body.SubscriberID,
	}
	v.BookID = bookID
	v.Token = token
	return v
}

// NewPickupPayload builds a books service pickup endpoint payload.
func NewPickupPayload(body *PickupRequestBody, bookID int64, token string) *books.PickupPayload {
	v := &books.PickupPayload{
		SubscriberID: *body.SubscriberID,
	}
	v.BookID = bookID
	v.Token = token
	return v
}

// NewReturnPayload builds a books service return endpoint payload.
func NewReturnPayload(body *ReturnRequestBody, token string) *books.ReturnPayload {
	v := &books.ReturnPayload{
		BookID:       *body.BookID,
		SubscriberID: *body.SubscriberID,
	}
	v.Token = token
	return v
}

// NewSubscribePayload builds a books service subscribe endpoint payload.
func NewSubscribePayload(bookID int64, token string) *books.SubscribePayload {
	return &books.SubscribePayload{
		BookID: bookID,
		Token:  token,
	}
}

// ValidateReserveRequestBody runs the validations defined on ReserveRequestBody
func ValidateReserveRequestBody(body *ReserveRequestBody) (err error) {
	if body.SubscriberID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("subscriber_id", "body"))
	}
	return
}

// ValidatePickupRequestBody runs the validations defined on PickupRequestBody
func ValidatePickupRequestBody(body *PickupRequestBody) (err error) {
	if body.SubscriberID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("subscriber_id", "body"))
	}
	return
}

// ValidateReturnRequestBody runs the validations defined on ReturnRequestBody
func ValidateReturnRequestBody(body *ReturnRequestBody) (err error) {
	if body.BookID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("book_id", "body"))
	}
	if body.SubscriberID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("subscriber_id", "body"))
	}
	return
}

// ValidateBookResponseBody runs the validations defined on BookResponseBody
func ValidateBookResponseBody(body *BookResponseBody) (err error) {
	if body.Images == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("images", "body"))
	}
	return
}
