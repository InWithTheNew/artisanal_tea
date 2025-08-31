package main

import (
	"fmt"
	"net/http"
	"os"

	"artisanal-kettle/controllers"
	"artisanal-kettle/internal/store"

	_ "artisanal-kettle/docs"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

func withCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		h(w, r)
	}
}

func main() {

	godotenv.Load()
	store.InitRedis((fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST"))), "", 0)

	r := mux.NewRouter()

	// User Handlers
	r.HandleFunc("/list/services", withCORS(controllers.ListServicesHandler)).Methods("GET", "OPTIONS")
	r.HandleFunc("/submit", withCORS(controllers.SubmitHandler)).Methods("POST", "OPTIONS")

	// Admin handlers
	r.HandleFunc("/admin/submit", withCORS(controllers.SubmitNewService)).Methods("POST")
	r.HandleFunc("/admin/delete", withCORS(controllers.DeleteService)).Methods("POST")

	// Docs
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusFound)
	})
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	http.ListenAndServe(":8080", r)
}
