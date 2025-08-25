package main

import (
	router "github/maskly15/go-backend/internal/router"
)

func main() {
	r := router.NewRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
