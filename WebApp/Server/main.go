package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/mskKandula/controller"
)

var (
	err error
)

func init() {
	controller.Db, err = sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/WebApp")

	if err != nil {
		log.Println("Connection Failed to Open")
	}
}

func main() {
	http.HandleFunc("/signup", controller.Signup)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	fmt.Println("server is running on 8080")
	http.ListenAndServe(":8080", nil)
}
