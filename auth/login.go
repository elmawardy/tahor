package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type SignInResponse struct {
	Jwt   string
	Name  string
	Email string
}

type SignInRequest struct {
	Email    string
	Password string
}

func GoSignIn(req SignInRequest) (res SignInResponse, err error) {

	u := &User{}
	err = u.ValidatePassword(req.Email, req.Password)
	if err != nil {
		return res, errors.New("Wrong username/password")
	}

	user := &User{}
	token, err := user.GenerateJWTFromEmail(req.Email)
	if err != nil {
		return res, err
	}

	res.Jwt = token
	return res, nil
}

func SignInHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req SignInRequest
		err := decodeRequestBody(r, &req)

		res, err := GoSignIn(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
		bytes, _ := json.Marshal(res)
		fmt.Fprintf(w, string(bytes))

		return
	}
}
