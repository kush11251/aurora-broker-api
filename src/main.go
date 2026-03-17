package main

import (
	"aurora-broker-api/src/controllers"
	"aurora-broker-api/src/models"
	"aurora-broker-api/src/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	controllers.InitializeRoutes()
	log.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
