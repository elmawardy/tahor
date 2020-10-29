package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RegisterRequest struct {
	// Display name of the user
	Name *string
	// Email used to perform signin
	Email string
	// Password used to perform signin
	Password string
}
type RegisterResponse struct {
	Message string
}

func Register(req RegisterRequest) (res RegisterResponse, err error) {
	u := &User{}
	u.Email = req.Email
	u.Password = req.Password
	if err != nil {
		return
	}
	u.Name = *req.Name

	err = u.Register()

	if err != nil {
		return
	}

	res.Message = "success"
	return
}

func RegisterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req RegisterRequest
		err := decodeRequestBody(r, &req)

		res, err := Register(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bytes, _ := json.Marshal(res)
		fmt.Fprintf(w, string(bytes))

		return
	}
}
