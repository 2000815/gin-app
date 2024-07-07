package model

import (
	"my-go-app/utility"

	"github.com/jinzhu/gorm"
)

type event struct {
	gorm.Model
	Id          int    `json:id`
	UserId     int    `json:user_id`
	Summary     string `json:summary`
	Dtstart     string `json:dtstart`
	Dtend       string `json:dtend`
	Description string `json:description`
	Year        int    `json:year`
	Month       int    `json:month`
	Day         int    `json:day`
}

func GetAll() (datas []event) {
	db := utility.Db
	result := db.Find(&datas)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}