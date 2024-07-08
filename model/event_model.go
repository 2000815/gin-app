package model

import (
	"github.com/jinzhu/gorm"
)

type Event struct {
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

// func GetAll(events []Event) (datas []Event) {

// 	result :=
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// 	return
// }