package main

import (
	"fmt"
	"net/http"

	"synthori.space/coffeeTime/internal/database"
	"synthori.space/coffeeTime/internal/routes"
)

func main() {
	fmt.Println("hello world")

	db := database.Connect()
	defer db.Close()

	r := routes.InitRoutes()

	http.ListenAndServe(":8000", r)
}
