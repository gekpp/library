// Code generated by goa v3.0.3, DO NOT EDIT.
//
// books client HTTP transport
//
// Command:
// $ goa gen library/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the books service endpoint HTTP clients.
type Client struct {
	// List Doer is the HTTP client used to make requests to the list endpoint.
	ListDoer goahttp.Doer

	// Reserve Doer is the HTTP client used to make requests to the reserve
	// endpoint.
	ReserveDoer goahttp.Doer

	// Pickup Doer is the HTTP client used to make requests to the pickup endpoint.
	PickupDoer goahttp.Doer

	// Return Doer is the HTTP client used to make requests to the return endpoint.
	ReturnDoer goahttp.Doer

	// Subscribe Doer is the HTTP client used to make requests to the subscribe
	// endpoint.
	SubscribeDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the books service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		ListDoer:            doer,
		ReserveDoer:         doer,
		PickupDoer:          doer,
		ReturnDoer:          doer,
		SubscribeDoer:       doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// List returns an endpoint that makes HTTP requests to the books service list
// server.
func (c *Client) List() goa.Endpoint {
	var (
		decodeResponse = DecodeListResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("books", "list", err)
		}
		return decodeResponse(resp)
	}
}

// Reserve returns an endpoint that makes HTTP requests to the books service
// reserve server.
func (c *Client) Reserve() goa.Endpoint {
	var (
		encodeRequest  = EncodeReserveRequest(c.encoder)
		decodeResponse = DecodeReserveResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildReserveRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ReserveDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("books", "reserve", err)
		}
		return decodeResponse(resp)
	}
}

// Pickup returns an endpoint that makes HTTP requests to the books service
// pickup server.
func (c *Client) Pickup() goa.Endpoint {
	var (
		encodeRequest  = EncodePickupRequest(c.encoder)
		decodeResponse = DecodePickupResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildPickupRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.PickupDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("books", "pickup", err)
		}
		return decodeResponse(resp)
	}
}

// Return returns an endpoint that makes HTTP requests to the books service
// return server.
func (c *Client) Return() goa.Endpoint {
	var (
		encodeRequest  = EncodeReturnRequest(c.encoder)
		decodeResponse = DecodeReturnResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildReturnRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ReturnDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("books", "return", err)
		}
		return decodeResponse(resp)
	}
}

// Subscribe returns an endpoint that makes HTTP requests to the books service
// subscribe server.
func (c *Client) Subscribe() goa.Endpoint {
	var (
		encodeRequest  = EncodeSubscribeRequest(c.encoder)
		decodeResponse = DecodeSubscribeResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSubscribeRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SubscribeDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("books", "subscribe", err)
		}
		return decodeResponse(resp)
	}
}
