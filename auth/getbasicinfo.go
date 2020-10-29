package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetBasicInfoRequest struct {
	JWT string
}
type GetBasicInfoResponse struct {
	Name string
}

func GetBasicInfo(req GetBasicInfoRequest) (res GetBasicInfoResponse, err error) {
	u := User{}
	info, err := u.GetBasicInfo(req.JWT)
	if err != nil {
		return
	}

	res.Name = info.Name
	return
}

func GetBasicInfoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req GetBasicInfoRequest
		err := decodeRequestBody(r, &req)

		res, err := GetBasicInfo(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
		bytes, _ := json.Marshal(res)
		fmt.Fprintf(w, string(bytes))

		return
	}
}
