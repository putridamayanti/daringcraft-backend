package models

import "time"

type Page struct {
	Id 				string			`json:"id"`
	Title 			string			`json:"title"`
	Code 			string			`json:"code"`
	Content 		string			`json:"content"`
	CreatedAt 		time.Time		`json:"createdAt"`
}
