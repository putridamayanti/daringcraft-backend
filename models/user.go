package models

import (
	"time"
)

type User struct {
	Id					string					`json:"id"`
	Active 				bool					`json:"active"`
	CreatedAt 			time.Time				`json:"createdAt"`
	CreatedBy			string					`json:"createdBy"`
	Email 				string					`json:"email"`
	Image				string					`json:"image"`
	LastActive			time.Time				`json:"lastActive"`
	Name 				string					`json:"name"`
	Password 			string					`json:"password"`
	Status 				bool					`json:"status"`
}

type Login struct {
	Email 		string		`json:"email" biding:"required"`
	Password 	string		`json:"password" binding:"required"`
}

type Register struct {
	Email 		string		`json:"email" biding:"required"`
	Name 		string		`json:"name" biding:"required"`
	Password 	string		`json:"password" biding:"required"`
}