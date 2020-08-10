package models

import (
	"github.com/elmawardy/tahor/global"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Case struct {
	gorm.Model
	Comment        string
	Collected      uint
	Target         uint
	Currency       string
	OrganizationID uint
	Votes          []User `gorm:"many2many:case_votes;"`
	Tags           []Tag
}

type CaseResponse struct {
	ID           uint
	Collected    uint
	Target       uint
	Currency     string
	Comment      string
	Tags         []string
	DateModified string
	OrgName      string
	OrgID        uint
}

func (c *Case) Insert() error {

	global.DB.Create(c)
	global.DB.Save(c)
	// global.DB.Model(c).Association("Tags").Append(c.Tags)
	return nil
}

func SelectCases(offset int, limit int) (cases []CaseResponse, e error) {

	// global.DB.Offset(offset).Limit(limit).Find(&cases)

	// var results interface{}
	// global.DB.Table("cases").Offset(offset).Limit(limit).Select("cases.id,cases.collected,cases.target,cases.comment,organizations.name,organizations.id").Joins("left join organizations on organizations.id = cases.organization_id").Scan(&cases)

	rows, _ := global.DB.Table("cases").Select("cases.id,cases.updated_at,cases.collected,cases.target,cases.currency,cases.comment,organizations.name,organizations.id").Joins("left join organizations on organizations.id = cases.organization_id").Offset(offset).Limit(limit).Rows()
	for rows.Next() {
		caseResponse := CaseResponse{}
		rows.Scan(&caseResponse.ID, &caseResponse.DateModified, &caseResponse.Collected, &caseResponse.Target, &caseResponse.Currency, &caseResponse.Comment, &caseResponse.OrgName, &caseResponse.OrgID)

		// global.DB.Table("Cases").Where("id = ?", caseResponse.ID).Association("Tags").Find(&caseResponse.Tags)
		tags := []Tag{}
		global.DB.Table("tags").Select("text").Where("case_id = ?", caseResponse.ID).Find(&tags)

		for _, tag := range tags {
			caseResponse.Tags = append(caseResponse.Tags, tag.Text)
		}

		cases = append(cases, caseResponse)

		// global.DB.Model()
	}

	return cases, nil
}
