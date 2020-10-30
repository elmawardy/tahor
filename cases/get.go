package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetCasesRequest struct {
	Token string
}
type GetCasesResponse struct {
	// Case ID
	ID *uint
	// Targetted value
	Target *float64
	// Collected value
	Collected *float64
	// Currency
	Currency *string
	// Case Description
	Description *string
	// Date Modified in the database
	DateModified *string
	// Tags
	Tags []string
}

func GetCases(req GetCasesRequest) (res []*GetCasesResponse, err error) {
	dbcases, err := SelectCases(0, 6)

	for i := 0; i < len(dbcases); i++ {
		response := GetCasesResponse{}
		response.Collected = &dbcases[i].Collected
		response.Currency = &dbcases[i].Currency
		response.DateModified = &dbcases[i].DateModified
		response.Description = &dbcases[i].Comment
		response.ID = &dbcases[i].ID
		response.Tags = dbcases[i].Tags
		response.Target = &dbcases[i].Target
		res = append(res, &response)
	}

	return
}

func GetCasesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req GetCasesRequest
		err := decodeRequestBody(r, &req)

		res, err := GetCases(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
		bytes, _ := json.Marshal(res)
		fmt.Fprintf(w, string(bytes))

		return
	}
}
