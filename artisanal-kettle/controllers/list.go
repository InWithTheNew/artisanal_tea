package controllers

import (
	"artisanal-kettle/internal/store"
	"encoding/json"
	"net/http"
)

// ListServicesHandler godoc
// @Summary      List all service names
// @Description  Returns a list of all registered service names
// @Tags         services
// @Produce      json
// @Success      200  {array}  string
// @Failure      500  {string}  string  "Error listing services"
// @Router       /list/services [get]
func ListServicesHandler(w http.ResponseWriter, r *http.Request) {
	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Retrieve all services from the store
	allServices, err := store.ListServices()
	var allServiceNames []string
	if err != nil {
		// If there was an error, return a 500 status and error message
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error listing services"))
	} else {
		// Loop through all services and append their names to the result slice
		for _, s := range allServices {
			allServiceNames = append(allServiceNames, s.GetName())
		}
		// Encode the list of service names as JSON and write to response
		json.NewEncoder(w).Encode(allServiceNames)
	}
}
