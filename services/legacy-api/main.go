package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	client, err := NewMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close(context.Background())

	apiHandler := NewAPIHandler(client)

	r := mux.NewRouter()

	r.HandleFunc("/rooms", apiHandler.CreateRoomHandler).Methods("POST")
	r.HandleFunc("/rooms/{id}", apiHandler.GetRoomByIDHandler).Methods("GET")
	r.HandleFunc("/rooms/{room_id}/tables", apiHandler.CreateTableHandler).Methods("POST")
	r.HandleFunc("/rooms/{room_id}/tables/{id}", apiHandler.GetTableByIDHandler).Methods("GET")
	r.HandleFunc("/rooms/{room_id}/tables/{id}", apiHandler.BookTable).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
