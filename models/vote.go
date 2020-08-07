package models

import "github.com/jinzhu/gorm"

type Vote struct {
	gorm.Model
	User     User
	Case     Case
	IsVoteUp bool
}
