package controllers

import (
	"artisanal-kettle/internal/service"
	"encoding/json"
	"io"
	"net/http"
)

// SubmitNewService godoc
// @Summary      Register a new service
// @Description  Registers a new service configuration in the system
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        service  body  service.Service  true  "Service definition"
// @Success      200  {string}  string  "OK"
// @Failure      400  {string}  string  "Invalid request"
// @Router       /admin/submit [post]
func SubmitNewService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s service.Service
	body, _ := io.ReadAll(r.Body)

	err := json.Unmarshal(body, &s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request"))
		return
	}

	err = s.SubmitNewServiceConfig()
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}

// SubmitNewService godoc
// @Summary      Deletes a new service
// @Description  Deletes a service configuration in the system
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        service  body  service.Service  true  "Service definition"
// @Success      200  {string}  string  "OK"
// @Failure      400  {string}  string  "Invalid request"
// @Router       /admin/submit [post]
func DeleteService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s service.Service
	body, _ := io.ReadAll(r.Body)

	err := json.Unmarshal(body, &s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request"))
		return
	}

	err = s.DeleteServiceConfig()
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}
