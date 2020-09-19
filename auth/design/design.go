package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("auth", func() {
	Title("Authentication")
	Description("API to authenticate users to be able to use the rest of services")
	Server("auth", func() {
		Host("localhost", func() { URI("http://localhost:8089") })
	})
})

var _ = Service("auth", func() {
	Description("A service to authenticate users")

	cors.Origin("*", func() {
		cors.Headers("Content-Type")
	})

	Method("register", func() {
		Description("Registers user in the database")

		Payload(func() {
			Description("Credentials used to authenticate to retrieve JWT token")
			Field(1, "name", String, "Display name of the user", func() {
				Example("Ahmad Sayed")
			})
			Field(2, "email", String, "Email used to perform signin", func() {
				Example("email.example.com")
			})
			Field(3, "password", String, "Password used to perform signin", func() {
				Example("password")
			})
			Required("email", "password")
		})

		Result(String)

		HTTP(func() {
			POST("/api/auth/register")
			// Use Authorization header to provide basic auth value.
			Response(StatusOK)
		})
	})

	Method("signin", func() {
		Description("Creates a valid JWT")

		Payload(func() {
			Description("Credentials used to authenticate to retrieve JWT token")
			Field(1, "email", String, "Email used to perform signin", func() {
				Example("email.example.com")
			})
			Field(2, "password", String, "Password used to perform signin", func() {
				Example("password")
			})
			Required("email", "password")
		})

		Result(Creds)

		HTTP(func() {
			POST("/api/auth/signin")
			// Use Authorization header to provide basic auth value.
			Response(StatusOK)
		})
	})

	Method("signout", func() {
		Description("Signs user out from the system")

		Payload(func() {
			Description("JWT used to authorize the process")
			Field(1, "jwt", String, "the user json web token", func() {
				Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
			})
			Required("jwt")
		})

		Result(String)

		HTTP(func() {
			POST("/api/auth/signout")
			// Use Authorization header to provide basic auth value.
			Response(StatusOK)
		})
	})

	Method("getbasicinfo", func() {
		Description("Gets user basic info like display name")

		Payload(func() {
			Description("JWT used to authorize the process")
			Field(1, "jwt", String, "the user json web token", func() {
				Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
			})
			Required("jwt")
		})

		Result(UserBasicInfo)

		HTTP(func() {
			POST("/api/auth/getbasicinfo")
			// Use Authorization header to provide basic auth value.
			Response(StatusOK)
		})
	})
})

// JWTAuth defines a security scheme that uses JWT tokens.
var JWTAuth = JWTSecurity("jwt", func() {
	Description(`Secures endpoint by requiring a valid JWT token retrieved via the signin endpoint. Supports scopes "api:read" and "api:write".`)
	Scope("api:read", "Read-only access")
	Scope("api:write", "Read and write access")
})

// Creds defines the credentials to use for authenticating to service methods.
var Creds = Type("Creds", func() {
	Field(1, "jwt", String, "JWT token", func() {
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
	})
	Required("jwt")
})

var UserBasicInfo = Type("UserBasicInfo", func() {
	Field(1, "name", String, "The user display name")
})
