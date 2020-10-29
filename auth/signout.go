package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SignOutRequest struct {
	JWT string
}
type SignOutResponse struct {
	Message string
}

func SignOut(req SignOutRequest) (res SignOutResponse, err error) {
	u := User{}
	err = u.Logout(req.JWT)
	res.Message = "success"
	return
}

func SignOutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req SignOutRequest
		err := decodeRequestBody(r, &req)

		res, err := SignOut(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
		bytes, _ := json.Marshal(res)
		fmt.Fprintf(w, string(bytes))

		return
	}
}
