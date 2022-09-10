package models

import "mime/multipart"

type Response struct {
	Status 		int				`json:"status"`
	Message 	string			`json:"message"`
	Data		interface{}		`json:"data"`
}

type File struct {
	File 		*multipart.FileHeader 		`form:"file"`
	MultiFile 	[]*multipart.FileHeader 	`form:"file"`
}

