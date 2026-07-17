package main

import (
	"fmt"
	"net/http"

	"synthori.space/coffeeTime/internal/routes"
)

func main() {
	fmt.Println("hello world")
	r := routes.InitRoutes()

	http.ListenAndServe(":8000", r)
}
