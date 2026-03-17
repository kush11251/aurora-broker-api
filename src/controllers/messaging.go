package controllers

import (
	"aurora-broker-api/src/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func InitializeRoutes() {
	http.HandleFunc("/messages", handleMessages)
}

func handleMessages(w http.ResponseWriter, r *http.Request) {
	var message models.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Message received: %+v", message)
}
