package design

import (
	goa "goa.design/goa/v3/dsl"
)

// BasicAuth defines a security scheme using basic authentication. The scheme
// protects the "signin" action used to create JWTs.
var BasicAuth = goa.BasicAuthSecurity("basic", func() {
	goa.Description("Basic authentication used to authenticate security principal during signin")
})

var _ = goa.Service("auther", func() {
	goa.Description("The auther service serves authentication methods")

	goa.HTTP()

	goa.Error("unauthorized", goa.String, "Credentials are invalid")

	goa.HTTP(func() {
		goa.Response("unauthorized", goa.StatusUnauthorized)
	})

	goa.Method("signin", func() {
		goa.Description("Creates a valid JWT")

		// The signin endpoint is secured via basic auth
		goa.Security(BasicAuth)

		goa.Payload(func() {
			goa.Description("Credentials used to authenticate to retrieve JWT token")
			goa.UsernameField(1, "username", goa.String, "Username used to perform signin", func() {
				goa.Example("user")
			})
			goa.PasswordField(2, "password", goa.String, "Password used to perform signin", func() {
				goa.Example("password")
			})
			goa.Required("username", "password")
		})

		goa.Result(Creds)

		goa.HTTP(func() {
			goa.POST("/signin")
			// Use Authorization header to provide basic auth value.
			goa.Response(goa.StatusOK)
		})
	})
})
