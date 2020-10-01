package handler

import "fmt"

// Signup is a endpoint handler to handle user registration
func Signup(username, password string) (err error) {
	// var uname string = "indra" other way to declare variable
	uname := "indra"
	pass := "indra123"
	if username != uname && password != pass {
		err = fmt.Errorf("invalid username or password")
		return
	}
	return
}
