package handlers

import (
	"encoding/json"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// Set the content type to be forced to use a json return type in the header
	w.Header().Set("Content-Type", "applications/json")

	// Set the response to a map where the key is the id of the json field and the value is the value
	response := map[string]string{"message": "Hi MICHELLE!!! THE API WORKS!!"}

	// json encoder encodes the map into a json object
	json.NewEncoder(w).Encode(response)
}
