package main

import (
	"fmt"
	"jardim-web/backend/internal/routes"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	routes.SetupRoutes(mux)
	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
