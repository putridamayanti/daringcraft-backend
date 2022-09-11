package models

import "time"

type Banner struct {
	Id 			string			`json:"id"`
	Title 		string			`json:"title"`
	Image 		string			`json:"image"`
	Status 		bool			`json:"status"`
	CreatedAt	time.Time		`json:"createdAt"`
}

type Mantra struct {
	Id 			string			`json:"id"`
	Mantra 		string			`json:"mantra"`
	Image 		string			`json:"image"`
	Status 		bool			`json:"status"`
	CreatedAt	time.Time		`json:"createdAt"`
}