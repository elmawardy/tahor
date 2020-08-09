package models

import (
	"github.com/elmawardy/tahor/global"
	"github.com/jinzhu/gorm"
)

type Vote struct {
	gorm.Model
	User     User
	Case     Case
	IsVoteUp bool
}

func (v *Vote) VoteUp(userid uint, caseid int) (err error) {

	v.IsVoteUp = true
	err = global.DB.First(&v.User).Where("id = ?", userid).Error
	if err != nil {
		return err
	}
	err = global.DB.First(&v.Case).Where("id = ?", caseid).Error
	if err != nil {
		return err
	}
	err = global.DB.Create(&v).Error
	if err != nil {
		return err
	}
	err = global.DB.Save(&v).Error
	if err != nil {
		return err
	}

	return err
}
