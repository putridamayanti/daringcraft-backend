package models

import "time"

type Subscriber struct {
	Id 				string			`json:"id"`
	Name 			string			`json:"name"`
	Email 			string			`json:"email"`
	CreatedAt		time.Time		`json:"createdAt"`
}
