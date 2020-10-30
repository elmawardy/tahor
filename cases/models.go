package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Case struct {
	gorm.Model
	Comment   string
	Collected float64
	Target    float64
	Currency  string
}

type CaseTag struct {
	gorm.Model
	CaseID uint
	Name   string
}

type CaseResponse struct {
	ID           uint
	Collected    float64
	Target       float64
	Currency     string
	Comment      string
	Tags         []string
	DateModified string
}

func (t *CaseTag) Insert() error {
	DB.Create(t)
	DB.Save(t)

	return nil
}

func (c *Case) Insert() (id uint, err error) {

	result := DB.Create(c)
	id = c.ID
	DB.Save(c)

	err = result.Error
	return
}

func SelectCases(offset int, limit int) (cases []CaseResponse, e error) {

	// global.DB.Offset(offset).Limit(limit).Find(&cases)

	// var results interface{}
	// global.DB.Table("cases").Offset(offset).Limit(limit).Select("cases.id,cases.collected,cases.target,cases.comment,organizations.name,organizations.id").Joins("left join organizations on organizations.id = cases.organization_id").Scan(&cases)

	rows, _ := DB.Table("cases").Select("cases.id,cases.updated_at,cases.collected,cases.target,cases.currency,cases.comment").Offset(offset).Limit(limit).Rows()
	for rows.Next() {
		caseResponse := CaseResponse{}
		rows.Scan(&caseResponse.ID, &caseResponse.DateModified, &caseResponse.Collected, &caseResponse.Target, &caseResponse.Currency, &caseResponse.Comment)

		// global.DB.Table("Cases").Where("id = ?", caseResponse.ID).Association("Tags").Find(&caseResponse.Tags)
		tags := []CaseTag{}
		DB.Table("case_tags").Select("name").Where("case_id = ?", caseResponse.ID).Find(&tags)

		for _, tag := range tags {
			caseResponse.Tags = append(caseResponse.Tags, tag.Name)
		}

		cases = append(cases, caseResponse)

		// global.DB.Model()
	}

	return cases, nil
}
