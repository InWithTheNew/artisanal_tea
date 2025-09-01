package controllers

import (
	"artisanal-kettle/internal/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// payload represents the expected structure of the JSON body for command submission.
type payload struct {
	Name    string // Name of the service to target
	Command string // Command to execute on the service
	User    string // User who has submitted the command
}

// SubmitHandler godoc
// @Summary      Submit a command to a service
// @Description  Submits a command to the specified service and returns the result
// @Tags         commands
// @Accept       json
// @Produce      json
// @Param        payload  body  controllers.payload  true  "Command payload"
// @Success      202  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /submit [post]
func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	var p payload
	// Read the entire request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// If reading the body fails, return a 400 error
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed to read body"})
		return
	}
	// Unmarshal the JSON body into the payload struct
	err = json.Unmarshal(body, &p)
	if err != nil {
		// If the JSON is invalid, return a 400 error
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid data struct submitted"})
		return
	}

	// Submit the command to the service layer
	response, err := service.SubmitCommand(p.Name, p.Command, p.User)
	if err != nil {
		// If the command fails, return a 500 error with details
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": fmt.Sprintf("Command rejected, %s", err)})
		return
	} else {
		// On success, return the result with a 202 Accepted status
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]string{"result": response})
	}
}
