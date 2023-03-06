package models

import (
	"github.com/adeindra6/skyshi-golang-test/app/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Activities struct {
	gorm.Model
	Title string `gorm:"column:title" json:"title"`
	Email string `gorm:"column:email" json:"email"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Activities{})
}

func (a *Activities) CreateActivities() *Activities {
	var result Activities

	db.Create(&a)
	db.Last(&result)

	return &result
}

func GetAllActitivities() []Activities {
	var Activities []Activities
	db.Find(&Activities)
	return Activities
}

func GetActivityById(id int64) (*Activities, *gorm.DB) {
	var getActivity Activities
	db := db.Where("id = ?", id).Find(&getActivity)
	return &getActivity, db
}

func DeleteActivity(id int64) bool {
	var activities Activities
	var findActivity int64

	db.Model(&Activities{}).Where("id = ?", id).Count(&findActivity)
	if findActivity > 0 {
		db.Model(&Activities{}).Where("id = ?", id).Delete(&activities)
		return true
	} else {
		return false
	}
}
