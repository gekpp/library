package booksapi

import (
	"context"
	auther "library/gen/auther"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// auther service example implementation.
// The example methods log the requests and return zero values.
type authersrvc struct {
	logger *log.Logger
}

// NewAuther returns the auther service implementation.
func NewAuther(logger *log.Logger) auther.Service {
	return &authersrvc{logger}
}

// Creates a valid JWT
func (s *authersrvc) Signin(ctx context.Context, p *auther.SigninPayload) (res *auther.Creds, err error) {
	// create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
			"iat":    time.Now().Unix(),
			"scopes": []string{"books:list", "books:reserve", "books:pickup", "books:return", "books:subscribe"},
			"sub":    p.Username,
		})

	s.logger.Printf("user '%s' signed in", p.Username)

	// note that if "SignedString" returns an error then it is returned as
	// an internal error to the client
	t, err := token.SignedString(Key)
	if err != nil {
		return nil, err
	}

	return &auther.Creds{
		JWT: t,
	}, nil
}
