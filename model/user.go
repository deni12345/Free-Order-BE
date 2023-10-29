package model

type User struct {
	UserName string "json: username"
	Age      int    "json: age"
	Email    string "json: email"
}
