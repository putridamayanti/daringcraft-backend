package models

type ForgotPasswordRequest struct {
	Email  		string			`json:"email"`
}

type UpdatePasswordRequest struct {
	Token 		string			`json:"token"`
	Password 	string			`json:"password"`
}

type UserRequest struct {
	Active 		bool			`json:"active"`
	Email 		string			`json:"email"`
	Main 		bool			`json:"main"`
	Name 		string			`json:"name"`
	Password 	string			`json:"password"`
	Status 		bool			`json:"status"`
}