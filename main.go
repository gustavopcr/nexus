package main

import (
	"fmt"
	"net/http"

	"github.com/gustavopcr/nexus/internal/routes"
)

func main() {
	router := routes.NewRouter()
	fmt.Println("starting server...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("error", err)
		return
	}
}
