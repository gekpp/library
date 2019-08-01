// Code generated by goa v3.0.3, DO NOT EDIT.
//
// books HTTP client CLI support package
//
// Command:
// $ goa gen library/design

package client

import (
	"encoding/json"
	"fmt"
	books "library/gen/books"
)

// BuildReservePayload builds the payload for the books reserve endpoint from
// CLI flags.
func BuildReservePayload(booksReserveBookID string, booksReserveToken string) (*books.ReservePayload, error) {
	var bookID string
	{
		bookID = booksReserveBookID
	}
	var token string
	{
		token = booksReserveToken
	}
	payload := &books.ReservePayload{
		BookID: bookID,
		Token:  token,
	}
	return payload, nil
}

// BuildPickupPayload builds the payload for the books pickup endpoint from CLI
// flags.
func BuildPickupPayload(booksPickupBody string, booksPickupBookID string, booksPickupToken string) (*books.PickupPayload, error) {
	var err error
	var body PickupRequestBody
	{
		err = json.Unmarshal([]byte(booksPickupBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"user_id\": \"Dolorum labore quas quidem sit fuga rerum.\"\n   }'")
		}
	}
	var bookID string
	{
		bookID = booksPickupBookID
	}
	var token string
	{
		token = booksPickupToken
	}
	v := &books.PickupPayload{
		UserID: body.UserID,
	}
	v.BookID = bookID
	v.Token = token
	return v, nil
}

// BuildReturnPayload builds the payload for the books return endpoint from CLI
// flags.
func BuildReturnPayload(booksReturnBody string, booksReturnBookID string, booksReturnToken string) (*books.ReturnPayload, error) {
	var err error
	var body ReturnRequestBody
	{
		err = json.Unmarshal([]byte(booksReturnBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"user_id\": \"Quo autem debitis exercitationem optio sed vitae.\"\n   }'")
		}
	}
	var bookID string
	{
		bookID = booksReturnBookID
	}
	var token string
	{
		token = booksReturnToken
	}
	v := &books.ReturnPayload{
		UserID: body.UserID,
	}
	v.BookID = bookID
	v.Token = token
	return v, nil
}

// BuildSubscribePayload builds the payload for the books subscribe endpoint
// from CLI flags.
func BuildSubscribePayload(booksSubscribeBookID string, booksSubscribeToken string) (*books.SubscribePayload, error) {
	var bookID string
	{
		bookID = booksSubscribeBookID
	}
	var token string
	{
		token = booksSubscribeToken
	}
	payload := &books.SubscribePayload{
		BookID: bookID,
		Token:  token,
	}
	return payload, nil
}
