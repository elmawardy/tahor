package models

import "github.com/jinzhu/gorm"

type Organization struct {
	gorm.Model
	Name  string
	Cases []Case
}
