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

// var Case = Type("Case", func() {
// 	// Attribute describes an object field
// 	Attribute("desc", String, "Case Description")
// 	Attribute("target", Float64, "Target Quantity")
// 	Attribute("collected", Float64, "Collected Quantity")
// 	Attribute("currency", String, "Currency")
// 	Attribute("tags", ArrayOf(String), "Tags array")

// })

var _ = Service("cases", func() {
	Description("Cases service")

	cors.Origin("*")

	Method("get", func() {
		Result(ArrayOf(GetResponse))
		HTTP(func() {
			GET("/api/cases/get")
			Response(StatusOK)
		})
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

// var sg = StorageGroup("PostgreSQL", func() {
// 	Description("This is the global storage group (relational)")
// 	Store("mysql", gorma.MySQL, func() {
// 		Description("This is the mysql relational store")
// 		Model("Case", func() {
// 			BuildsFrom()

// 			RendersTo(Cases)						// a Media Type definition
// 			Description("This is the case model")

// 			Field("ID", gorma.Integer, func() {    // Required for CRUD getters to take a PK argument!
// 				PrimaryKey()
// 				Description("This is the ID PK field")
// 			})

// 			Field("Comment", gorma.String)
// 			Field("Target", gorma.Float64)
// 			Field("Collected", gorma.Float64)
// 			Field("Tags", gorma.ArrayOf(String))

// 			Field("CreatedAt", gorma.Timestamp)
// 			Field("UpdatedAt", gorma.Timestamp)			 // Shown for demonstration
// 			Field("DeletedAt", gorma.NullableTimestamp)  // These are added by default
// 		})
// 	})
// }

// var _ = Service("openapi", func() {
// 	// Serve the file with relative path ../../gen/http/openapi.json for
// 	// requests sent to /swagger.json.
// 	Files("/swagger.json", "../../gen/http/openapi.json")
// })
