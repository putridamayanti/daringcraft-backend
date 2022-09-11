package models

import (
	"mime/multipart"
	"time"
)

type Response struct {
	Status 		int				`json:"status"`
	Message 	string			`json:"message"`
	Data		interface{}		`json:"data"`
}

type File struct {
	File 		*multipart.FileHeader 		`form:"file"`
	MultiFile 	[]*multipart.FileHeader 	`form:"file"`
}

type Media struct {
	Id 			string			`json:"id"`
	Url 		*string			`json:"url"`
	CreatedAt	time.Time		`json:"createdAt"`
}