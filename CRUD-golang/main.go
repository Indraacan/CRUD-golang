package main

import (
	"fmt"
	"os"

	"github.com/Indraacan/CRUD-golang/handler"
	"github.com/Indraacan/CRUD-golang/model"
)

func main() {
	user := &model.User{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
	err := handler.Signin(user.Username, user.Password)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success login")
}
