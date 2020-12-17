package main

import (
	db "BGP-Golang-TRPL3A/echo-rest/db"
	routes "BGP-Golang-TRPL3A/echo-rest/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":3000"))
}
