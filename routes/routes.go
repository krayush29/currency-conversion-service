package routes

import (
	"currency-conversion-service/controllers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/convert", controllers.ConvertHandler)
}
