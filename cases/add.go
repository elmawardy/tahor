package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AddCaseRequest struct {
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
type AddCaseResponse struct {
	Desc *string
}

func AddCase(req AddCaseRequest) (res AddCaseResponse, err error) {
	res.Desc = &req.Desc

	newCase := &Case{}
	newCase.Comment = req.Desc
	newCase.Collected = req.Collected
	newCase.Target = req.Target
	newCase.Currency = req.Currency
	caseid, err := newCase.Insert()

	if req.Tags != nil {
		for i := 0; i < len(req.Tags); i++ {
			tag := &CaseTag{}
			tag.CaseID = caseid
			tag.Name = req.Tags[i]
			tag.Insert()
		}
	}

	return
}

func AddCaseHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req AddCaseRequest
		err := decodeRequestBody(r, &req)

		res, err := AddCase(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
		bytes, _ := json.Marshal(res)
		fmt.Fprintf(w, string(bytes))

		return
	}
}
