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

// PickupRequestBody is the type of the "books" service "pickup" endpoint HTTP
// request body.
type PickupRequestBody struct {
	// id of the user picking up the book
	UserID *string `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
}

// ReturnRequestBody is the type of the "books" service "return" endpoint HTTP
// request body.
type ReturnRequestBody struct {
	// id of the user returning the book
	UserID *string `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
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
	ID         string `form:"id" json:"id" xml:"id"`
	Title      string `form:"title" json:"title" xml:"title"`
	Annotation string `form:"annotation" json:"annotation" xml:"annotation"`
	Author     string `form:"author" json:"author" xml:"author"`
	// images are a list of book photos
	Images []string `form:"images" json:"images" xml:"images"`
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
func NewReservePayload(bookID string, token string) *books.ReservePayload {
	return &books.ReservePayload{
		BookID: bookID,
		Token:  token,
	}
}

// NewPickupPayload builds a books service pickup endpoint payload.
func NewPickupPayload(body *PickupRequestBody, bookID string, token string) *books.PickupPayload {
	v := &books.PickupPayload{
		UserID: *body.UserID,
	}
	v.BookID = bookID
	v.Token = token
	return v
}

// NewReturnPayload builds a books service return endpoint payload.
func NewReturnPayload(body *ReturnRequestBody, bookID string, token string) *books.ReturnPayload {
	v := &books.ReturnPayload{
		UserID: *body.UserID,
	}
	v.BookID = bookID
	v.Token = token
	return v
}

// NewSubscribePayload builds a books service subscribe endpoint payload.
func NewSubscribePayload(bookID string, token string) *books.SubscribePayload {
	return &books.SubscribePayload{
		BookID: bookID,
		Token:  token,
	}
}

// ValidatePickupRequestBody runs the validations defined on PickupRequestBody
func ValidatePickupRequestBody(body *PickupRequestBody) (err error) {
	if body.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user_id", "body"))
	}
	return
}

// ValidateReturnRequestBody runs the validations defined on ReturnRequestBody
func ValidateReturnRequestBody(body *ReturnRequestBody) (err error) {
	if body.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user_id", "body"))
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