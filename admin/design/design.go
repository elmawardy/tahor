package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("admin", func() {
	Title("Admininstration app")
	Description("API to administrate the project like making crud for cases")
	Server("admin", func() {
		Host("localhost", func() { URI("http://localhost:8090") })
	})
})

var _ = Service("admin", func() {
	Description("A service to administrate the project")

	Security(JWTAuth)

	cors.Origin("*", func() {
		cors.Headers("Content-Type")
	})

	Method("addcase", func() {
		// Payload describes the method payload
		// Here the payload is an object that consists of two fields
		Payload(func() {
			// Attribute describes an object field
			Token("jwt", String, "JWT token used to perform authorization")
			Attribute("desc", String, "Case Description")
			Attribute("target", Float64, "Target Quantity")
			Attribute("collected", Float64, "Collected Quantity")
			Attribute("currency", String, "Currency")
			Attribute("tags", ArrayOf(String), "Tags array")
			Required("desc", "target", "collected", "currency")
		})
		// Result describes the method result
		// Here the result is a simple integer value
		Result(String)
		// HTTP describes the HTTP transport mapping
		HTTP(func() {
			// Requests to the service consist of HTTP GET requests
			// The payload fields are encoded as path parameters
			POST("/api/admin/cases/add")
			// Responses use a "200 OK" HTTP status
			// The result is encoded in the response body
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
