package casesapi

import (
	"context"
	"log"

	cases "github.com/elmawardy/tahor/cases/gen/cases"
	"github.com/elmawardy/tahor/cases/models"
)

// cases service example implementation.
// The example methods log the requests and return zero values.
type casessrvc struct {
	logger *log.Logger
}

// NewCases returns the cases service implementation.
func NewCases(logger *log.Logger) cases.Service {
	return &casessrvc{logger}
}

// Add implements add.
func (s *casessrvc) Add(ctx context.Context, p *cases.AddPayload) (res *cases.AddResponse, err error) {
	res = &cases.AddResponse{}
	s.logger.Print("cases.add")
	res.Desc = &p.Desc

	newCase := &models.Case{}
	newCase.Comment = p.Desc
	newCase.Collected = p.Collected
	newCase.Target = p.Target
	newCase.Currency = p.Currency
	caseid, err := newCase.Insert()

	if p.Tags != nil {
		for i := 0; i < len(p.Tags); i++ {
			tag := &models.CaseTag{}
			tag.CaseID = caseid
			tag.Name = p.Tags[i]
			tag.Insert()
		}
	}

	return
}

func (s *casessrvc) Get(ctx context.Context, p *cases.GetPayload) (res []*cases.GetResponse, err error) {

	dbcases, err := models.SelectCases(0, 6)

	for i := 0; i < len(dbcases); i++ {
		response := cases.GetResponse{}
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
