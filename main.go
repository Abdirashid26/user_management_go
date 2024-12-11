package main

import (
	"fmt"
	"log"
	"net/http"
	"user_management_go/routes"
)

func main() {
	fmt.Println("Server is running on port 8080 .....")
	log.Fatal(http.ListenAndServe("localhost:8080", routes.SetupRoutes()))
}
