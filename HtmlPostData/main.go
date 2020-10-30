package main

import (
	"fmt"
	"log"
	"net/http"
	fn "BGP-Golang-TRPL3A/HtmlPostData/function"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
    http.HandleFunc("/", fn.RouteIndexGet)
    http.HandleFunc("/process", fn.RouteSubmitPost)

		fmt.Println("server started at localhost:9000")
		log.Println("Server started on: http://localhost:9000/")
    http.ListenAndServe(":9000", nil)
}



