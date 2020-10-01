// Code generated by goa v3.2.4, DO NOT EDIT.
//
// admin service
//
// Command:
// $ goa gen github.com/elmawardy/tahor/admin/design

package admin

import (
	"context"

	"goa.design/goa/v3/security"
)

// A service to administrate the project
type Service interface {
	// Addcase implements addcase.
	Addcase(context.Context, *AddcasePayload) (res string, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "admin"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"addcase"}

// AddcasePayload is the payload type of the admin service addcase method.
type AddcasePayload struct {
	// JWT token used to perform authorization
	JWT *string
	// Case Description
	Desc string
	// Target Quantity
	Target float64
	// Collected Quantity
	Collected float64
	// Currency
	Currency string
	// Tags array
	Tags []string
}
