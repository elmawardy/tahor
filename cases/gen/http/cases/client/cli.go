// Code generated by goa v3.2.3, DO NOT EDIT.
//
// cases HTTP client CLI support package
//
// Command:
// $ goa gen github.com/elmawardy/tahor/cases/design

package client

import (
	"encoding/json"
	"fmt"

	cases "github.com/elmawardy/tahor/cases/gen/cases"
)

// BuildGetPayload builds the payload for the cases get endpoint from CLI flags.
func BuildGetPayload(casesGetBody string) (*cases.GetPayload, error) {
	var err error
	var body GetRequestBody
	{
		err = json.Unmarshal([]byte(casesGetBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"token\": \"Est esse id.\"\n   }'")
		}
	}
	v := &cases.GetPayload{
		Token: body.Token,
	}

	return v, nil
}

// BuildAddPayload builds the payload for the cases add endpoint from CLI flags.
func BuildAddPayload(casesAddBody string) (*cases.AddPayload, error) {
	var err error
	var body AddRequestBody
	{
		err = json.Unmarshal([]byte(casesAddBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"collected\": 0.8003833617014785,\n      \"currency\": \"Sunt libero ullam nam voluptatem quis.\",\n      \"desc\": \"Qui officia non ipsa accusantium facilis consectetur.\",\n      \"tags\": [\n         \"Quisquam vel deserunt.\",\n         \"Enim esse voluptas.\",\n         \"Ut quia ratione ipsa repellat.\"\n      ],\n      \"target\": 0.34487022127545675\n   }'")
		}
	}
	v := &cases.AddPayload{
		Desc:      body.Desc,
		Target:    body.Target,
		Collected: body.Collected,
		Currency:  body.Currency,
	}
	if body.Tags != nil {
		v.Tags = make([]string, len(body.Tags))
		for i, val := range body.Tags {
			v.Tags[i] = val
		}
	}

	return v, nil
}
