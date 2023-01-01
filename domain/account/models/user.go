package models

type User struct {
	Id       string `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}
