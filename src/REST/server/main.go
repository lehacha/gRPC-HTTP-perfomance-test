package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

type ProcessRequest struct {
	Text     string `json:"text"`
	Username string `json:"username"`
}

type ProcessResponse struct {
	Words   []string `json:"words"`
	Message string   `json:"message"`
}

func main() {
	log.Println("[REST] Starting server")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.AllowAll().Handler)

	r.Post("/process", func(w http.ResponseWriter, r *http.Request) {

		// Declare a new ProcessRequest struct.
		var processRequest ProcessRequest

		// Try to decode the request body into the struct
		err := json.NewDecoder(r.Body).Decode(&processRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		log.Printf("[REST] Receive message body from client: %+v", processRequest)

		var respB []byte
		respB, err = json.Marshal(ProcessResponse{
			Words:   []string{"test"},
			Message: "test",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		_, err = w.Write(respB)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
		w.WriteHeader(http.StatusOK)
	})

	err := http.ListenAndServe(":8090", r)
	if err != nil {
		log.Panic("Can not start mil service: " + err.Error())
	}
}
