package main

import (
	"currency-conversion-service/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	routes.RegisterRoutes()
	fmt.Println("Currency conversion service running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
