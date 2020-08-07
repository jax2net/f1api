package main

import (
	"f1api/routes"
	"fmt"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on port 8081...")
	http.ListenAndServe(":8081", r)
}
