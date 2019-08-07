// Code generated by goa v3.0.3, DO NOT EDIT.
//
// books HTTP server
//
// Command:
// $ goa gen library/design

package server

import (
	"context"
	books "library/gen/books"
	"net/http"
	"regexp"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the books service endpoint HTTP handlers.
type Server struct {
	Mounts    []*MountPoint
	List      http.Handler
	Reserve   http.Handler
	Pickup    http.Handler
	Return    http.Handler
	Subscribe http.Handler
	CORS      http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the books service endpoints.
func New(
	e *books.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"List", "GET", "/books/list"},
			{"Reserve", "POST", "/books/reserve/{book_id}"},
			{"Pickup", "POST", "/books/pickup/{book_id}"},
			{"Return", "POST", "/books/return"},
			{"Subscribe", "POST", "/books/subscribe/{book_id}"},
			{"CORS", "OPTIONS", "/books/list"},
			{"CORS", "OPTIONS", "/books/reserve/{book_id}"},
			{"CORS", "OPTIONS", "/books/pickup/{book_id}"},
			{"CORS", "OPTIONS", "/books/return"},
			{"CORS", "OPTIONS", "/books/subscribe/{book_id}"},
		},
		List:      NewListHandler(e.List, mux, dec, enc, eh),
		Reserve:   NewReserveHandler(e.Reserve, mux, dec, enc, eh),
		Pickup:    NewPickupHandler(e.Pickup, mux, dec, enc, eh),
		Return:    NewReturnHandler(e.Return, mux, dec, enc, eh),
		Subscribe: NewSubscribeHandler(e.Subscribe, mux, dec, enc, eh),
		CORS:      NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "books" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.List = m(s.List)
	s.Reserve = m(s.Reserve)
	s.Pickup = m(s.Pickup)
	s.Return = m(s.Return)
	s.Subscribe = m(s.Subscribe)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the books endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountListHandler(mux, h.List)
	MountReserveHandler(mux, h.Reserve)
	MountPickupHandler(mux, h.Pickup)
	MountReturnHandler(mux, h.Return)
	MountSubscribeHandler(mux, h.Subscribe)
	MountCORSHandler(mux, h.CORS)
}

// MountListHandler configures the mux to serve the "books" service "list"
// endpoint.
func MountListHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleBooksOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/books/list", f)
}

// NewListHandler creates a HTTP handler which loads the HTTP request and calls
// the "books" service "list" endpoint.
func NewListHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		encodeResponse = EncodeListResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "list")
		ctx = context.WithValue(ctx, goa.ServiceKey, "books")

		res, err := endpoint(ctx, nil)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountReserveHandler configures the mux to serve the "books" service
// "reserve" endpoint.
func MountReserveHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleBooksOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/books/reserve/{book_id}", f)
}

// NewReserveHandler creates a HTTP handler which loads the HTTP request and
// calls the "books" service "reserve" endpoint.
func NewReserveHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeReserveRequest(mux, dec)
		encodeResponse = EncodeReserveResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "reserve")
		ctx = context.WithValue(ctx, goa.ServiceKey, "books")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountPickupHandler configures the mux to serve the "books" service "pickup"
// endpoint.
func MountPickupHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleBooksOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/books/pickup/{book_id}", f)
}

// NewPickupHandler creates a HTTP handler which loads the HTTP request and
// calls the "books" service "pickup" endpoint.
func NewPickupHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodePickupRequest(mux, dec)
		encodeResponse = EncodePickupResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "pickup")
		ctx = context.WithValue(ctx, goa.ServiceKey, "books")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountReturnHandler configures the mux to serve the "books" service "return"
// endpoint.
func MountReturnHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleBooksOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/books/return", f)
}

// NewReturnHandler creates a HTTP handler which loads the HTTP request and
// calls the "books" service "return" endpoint.
func NewReturnHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeReturnRequest(mux, dec)
		encodeResponse = EncodeReturnResponse(enc)
		encodeError    = EncodeReturnError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "return")
		ctx = context.WithValue(ctx, goa.ServiceKey, "books")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountSubscribeHandler configures the mux to serve the "books" service
// "subscribe" endpoint.
func MountSubscribeHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleBooksOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/books/subscribe/{book_id}", f)
}

// NewSubscribeHandler creates a HTTP handler which loads the HTTP request and
// calls the "books" service "subscribe" endpoint.
func NewSubscribeHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeSubscribeRequest(mux, dec)
		encodeResponse = EncodeSubscribeResponse(enc)
		encodeError    = EncodeSubscribeError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "subscribe")
		ctx = context.WithValue(ctx, goa.ServiceKey, "books")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service books.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = handleBooksOrigin(h)
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("OPTIONS", "/books/list", f)
	mux.Handle("OPTIONS", "/books/reserve/{book_id}", f)
	mux.Handle("OPTIONS", "/books/pickup/{book_id}", f)
	mux.Handle("OPTIONS", "/books/return", f)
	mux.Handle("OPTIONS", "/books/subscribe/{book_id}", f)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// handleBooksOrigin applies the CORS response headers corresponding to the
// origin for the service books.
func handleBooksOrigin(h http.Handler) http.Handler {
	spec0 := regexp.MustCompile(".*localhost.*")
	origHndlr := h.(http.HandlerFunc)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			origHndlr(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec0) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "X-Time, X-Api-Version")
			w.Header().Set("Access-Control-Max-Age", "100")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
				w.Header().Set("Access-Control-Allow-Headers", "X-Shared-Secret")
			}
			origHndlr(w, r)
			return
		}
		origHndlr(w, r)
		return
	})
}
