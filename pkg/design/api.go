package design

import (
	goa "goa.design/goa/v3/dsl"
)

// API describes the global properties of the API server.
var _ = goa.API("books", func() {
	goa.Title("Books Service")
	goa.Description("HTTP service for working with books")
	goa.Server("books", func() {
		goa.Services("books", "auther")
		goa.Host("localhost", func() { goa.URI("http://localhost:8088") })
	})
})

// Book is a book
var Book = goa.Type("Book", func() {
	goa.Description("Book is a book")

	goa.Attribute("id", goa.Int64)
	goa.Attribute("title", goa.String)
	goa.Attribute("annotation", goa.String)
	goa.Attribute("author", goa.String)
	goa.Attribute("images", goa.ArrayOf(goa.String), func() {
		goa.Description("images are a list of book photos")
	})
	goa.Attribute("status", goa.String)

	goa.Required("id", "title", "annotation", "author", "images", "status")
})

// ListOfBooks is a List of books
var ListOfBooks = goa.ResultType("application/library.books", func() {
	goa.Description("List of books")

	goa.Attribute("data", goa.ArrayOf(Book))

	goa.Required("data")
})

// Creds defines the credentials to use for authenticating to service methods.
var Creds = goa.Type("Creds", func() {
	goa.Field(1, "jwt", goa.String, "JWT token", func() {
		goa.Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
	})
	goa.Required("jwt")
})
