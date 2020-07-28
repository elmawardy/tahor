package models

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Text   string
	CaseID uint
}
