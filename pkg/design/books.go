package design

import (
	goa "goa.design/goa/v3/dsl"
)

// JWTAuth defines a security scheme that uses JWT tokens.
var JWTAuth = goa.JWTSecurity("jwt", func() {
	goa.Description(`Secures endpoint by requiring a valid JWT token retrieved via the signin endpoint. 
	Supports scopes "books:reserve", "books:pickup", "books:return", "books:subscribe"`)
	goa.Scope("books:reserve", "reserve access")
	goa.Scope("books:pickup", "pickup access")
	goa.Scope("books:return", "return access")
	goa.Scope("books:subscribe", "subscribe access")
})

var _ = goa.Service("books", func() {
	goa.Description("The books service serves operations on books: list, reserve, pickedUp, returned, subscribe")

	goa.HTTP(func() {
		goa.Path("/books")
	})

	goa.Method("list", func() {
		goa.Description("List books")
		goa.Result(ListOfBooks)
		goa.HTTP(func() {
			goa.GET("/list")
			goa.Response(goa.StatusOK)
		})
	})

	goa.Method("reserve", func() {
		goa.Description(`Mark book as reserved. Once a book is reserved timer starts with timeout for the book to become picked up. Timeout is configurable. 
		Once timeout is expired book becomes available`)
		goa.Result(goa.Any)

		jwtScopeSecurity("books:reserve")

		goa.Payload(func() {
			goa.TokenField(0, "token", goa.String, func() {
				goa.Description("JWT used for authentication")
			})
			goa.Field(1, "book_id", goa.String, func() {
				goa.Description("id of the Book")
			})

			goa.Required("token", "book_id")
		})

		goa.HTTP(func() {
			goa.POST("/reserve/{book_id}")
			goa.Response(goa.StatusOK)
		})
	})

	goa.Method("pickup", func() {
		goa.Description("Mark book as picked up")
		goa.Result(goa.Any)

		jwtScopeSecurity("books:pickup")

		goa.Payload(func() {
			goa.TokenField(0, "token", goa.String, func() {
				goa.Description("JWT used for authentication")
			})
			goa.Field(1, "book_id", goa.String, func() {
				goa.Description("id of the Book")
			})
			goa.Field(2, "user_id", goa.String, func() {
				goa.Description("id of the user picking up the book")
			})

			goa.Required("token", "book_id", "user_id")
		})

		goa.HTTP(func() {
			goa.POST("/pickup/{book_id}")
			goa.Response(goa.StatusOK)
		})
	})

	goa.Method("return", func() {
		goa.Description("Mark book as returned")
		goa.Result(goa.Any)

		jwtScopeSecurity("books:return")

		goa.Payload(func() {
			goa.TokenField(0, "token", goa.String, func() {
				goa.Description("JWT used for authentication")
			})

			goa.Field(1, "book_id", goa.String, func() {
				goa.Description("id of the Book")
			})
			goa.Field(2, "user_id", goa.String, func() {
				goa.Description("id of the user returning the book")
			})

			goa.Required("token", "book_id", "user_id")
		})

		goa.HTTP(func() {
			goa.POST("/return/{book_id}")
			goa.Response(goa.StatusOK)
			goa.Response("invalid-scopes", goa.StatusForbidden)
		})
	})

	goa.Method("subscribe", func() {
		goa.Description("Subscribe the caller on the next 'book's become available")
		goa.Result(goa.Any)

		jwtScopeSecurity("books:subscribe")

		goa.Payload(func() {
			goa.TokenField(0, "token", goa.String, func() {
				goa.Description("JWT used for authentication")
			})
			goa.Field(1, "book_id", goa.String, func() {
				goa.Description("id of the Book")
			})

			goa.Required("token", "book_id")
		})

		goa.HTTP(func() {
			goa.POST("/subscribe/{book_id}")
			goa.Response(goa.StatusOK)
			goa.Response("invalid-scopes", goa.StatusForbidden)
		})
	})
})

func jwtScopeSecurity(scopes ...string) {
	goa.Security(JWTAuth, func() { // Use JWT and an API key to secure this endpoint.
		for _, scope := range scopes {
			goa.Scope(scope)
		}
	})
	goa.Error("invalid-scopes", goa.String, "Token scopes are invalid")
}
