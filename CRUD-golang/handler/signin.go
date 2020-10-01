package handler

import "fmt"

// Signin is a endpoint handler to handle user login
func Signin(username, password string) (err error) {
	// var uname string = "indra" other way to declare variable
	uname := "indra"
	pass := "indra123"
	if username != uname && password != pass {
		err = fmt.Errorf("invalid username or password")
		return
	}
	return
}
