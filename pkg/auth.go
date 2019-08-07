package booksapi

import (
	"context"
	"library/gen/auther"
	books "library/gen/books"

	"github.com/google/uuid"

	"github.com/dgrijalva/jwt-go"
	"goa.design/goa/v3/security"
)

var (
	// ErrUnauthorized is the error returned by Login when the request credentials
	// are invalid.
	ErrUnauthorized error = auther.Unauthorized("invalid username and password combination")

	// ErrInvalidToken is the error returned when the JWT token is invalid.
	ErrInvalidToken error = auther.Unauthorized("invalid token")

	// ErrInvalidTokenScopes is the error returned when the scopes provided in
	// the JWT token claims are invalid.
	ErrInvalidTokenScopes error = books.InvalidScopes("invalid scopes in token")

	// Key is the key used in JWT authentication
	Key = []byte("secret")

	// HardcodedUserID is the one-for-all (very unsecure!!!) ID of the service caller if applicable.
	HardcodedUserID, _ = uuid.NewRandom()
)

type contextKey string

var (
	contextKeyUserID = contextKey("username")
)

// BasicAuth implements the authorization logic for service "auther" for the
// "basic" security scheme.
func (s *authersrvc) BasicAuth(ctx context.Context, user, pass string, scheme *security.BasicScheme) (context.Context, error) {
	if user == "librarian" && pass == "library" {
		return ctx, nil
	}

	return ctx, ErrUnauthorized
}

// JWTAuth implements the authorization logic for service "books" for the "jwt"
// security scheme.
func (s *bookssrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	claims := make(jwt.MapClaims)

	// authorize request
	// 1. parse JWT token, token key is hardcoded to "secret" in this example
	_, err := jwt.ParseWithClaims(token, claims, func(_ *jwt.Token) (interface{}, error) { return Key, nil })
	if err != nil {
		return ctx, ErrInvalidToken
	}

	// 2. validate provided "scopes" claim
	if claims["scopes"] == nil {
		return ctx, ErrInvalidTokenScopes
	}
	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		return ctx, ErrInvalidTokenScopes
	}
	scopesInToken := make([]string, 0, len(scopes))
	for _, scp := range scopes {
		scopesInToken = append(scopesInToken, scp.(string))
	}
	if err := scheme.Validate(scopesInToken); err != nil {
		return ctx, books.InvalidScopes(err.Error())
	}

	userID, err := uuid.Parse(claims["sub"].(string))
	if err != nil {
		return ctx, books.InvalidScopes("Unparsable userID")
	}

	return context.WithValue(ctx, contextKeyUserID, userID), nil
}
