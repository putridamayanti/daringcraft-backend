package models

import "time"

type Message struct {
	Id 				string			`json:"id"`
	Name 			string			`json:"name"`
	Email 			string			`json:"email"`
	Message			string			`json:"message"`
	CreatedAt		time.Time		`json:"createdAt"`
}

type Customer struct {
	Id 				string			`json:"id"`
	Name 			string			`json:"name"`
	Email 			string			`json:"email"`
	CreatedAt		time.Time		`json:"createdAt"`
}