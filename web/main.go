package main

import "github.com/alberliu/goweb"

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func handler(user User) User {
	return user
}

func main() {
	goweb.HandlePost("/test", handler)
	goweb.ListenAndServe(":8000")
}
