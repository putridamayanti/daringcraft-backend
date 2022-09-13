package models

import "time"

type Post struct {
	Id 				string			`json:"id"`
	Title 			string			`json:"title"`
	Content 		string			`json:"content"`
	Tags 			[]string		`json:"tags"`
	Slug 			string			`json:"slug"`
	Status 			string			`json:"status"` // DRAFT, PUBLISHED, ARCHIVED
	CreatedAt 		time.Time		`json:"createdAt"`
}
