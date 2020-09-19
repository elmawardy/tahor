package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("cases", func() {
	Title("Cases")
	Description("API to make CRUD operations for the cases")
	Server("cases", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})

var _ = Service("cases", func() {
	Description("Cases service")

	cors.Origin("*")

	Error("unauthorized", String, "Credentials are invalid")

	HTTP(func() {
		Response("unauthorized", StatusUnauthorized)
	})

	Method("get", func() {
		Result(ArrayOf(GetResponse))
		Security(JWTAuth, func() { // Use JWT to auth requests to this endpoint.
			Scope("api:read") // Enforce presence of "api:read" scope in JWT claims.
		})
		Payload(func() {
			TokenField(1, "token", String, func() {
				Description("JWT used for authentication")
			})
			Required("token")
		})
		HTTP(func() {
			GET("/api/cases/get")
			Response(StatusOK)
			Response("invalid-scopes", StatusForbidden)
		})
		Error("invalid-scopes", String, "Token scopes are invalid")
	})
	Method("add", func() {
		// Payload describes the method payload
		// Here the payload is an object that consists of two fields
		Payload(func() {
			// Attribute describes an object field
			Attribute("desc", String, "Case Description")
			Attribute("target", Float64, "Target Quantity")
			Attribute("collected", Float64, "Collected Quantity")
			Attribute("currency", String, "Currency")
			Attribute("tags", ArrayOf(String), "Tags array")
			Required("desc", "target", "collected", "currency")
		})
		// Result describes the method result
		// Here the result is a simple integer value
		Result(AddResponse)
		// HTTP describes the HTTP transport mapping
		HTTP(func() {
			// Requests to the service consist of HTTP GET requests
			// The payload fields are encoded as path parameters
			POST("/api/cases/add")
			// Responses use a "200 OK" HTTP status
			// The result is encoded in the response body
			Response(StatusOK)
		})

	})
})

var AddResponse = Type("AddResponse", func() {
	Attribute("desc", String, "The description of the case")
})

var GetResponse = Type("GetResponse", func() {
	Attribute("ID", UInt, "Case ID")
	Attribute("Target", Float64, "Targetted value")
	Attribute("Collected", Float64, "Collected value")
	Attribute("Currency", String, "Currency")
	Attribute("Description", String, "Case Description")
	Attribute("DateModified", String, "Date Modified in the database")
	Attribute("Tags", ArrayOf(String), "Tags")
})
