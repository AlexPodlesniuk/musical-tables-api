package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Room struct {
	ID        string             `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Tables    []Table            `json:"tables,omitempty" bson:"tables,omitempty"`
	CreatedAt primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type Table struct {
	ID        string             `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Capacity  int                `json:"capacity" bson:"capacity"`
	IsBooked  bool               `json:"isBooked" bson:"isBooked"`
	CreatedAt primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type APIHandler struct {
	client *MongoClient
}

func NewAPIHandler(client *MongoClient) *APIHandler {
	return &APIHandler{client: client}
}

func (h *APIHandler) CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	var room Room
	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	room.ID = uuid.New().String()
	err = h.client.SaveRoom(r.Context(), &room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(room)
}

func (h *APIHandler) GetRoomByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	roomID := params["id"]

	room, err := h.client.GetRoomByID(r.Context(), roomID)
	if err != nil {
		if errors.Is(err, ErrRoomNotFound) {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(room)
}

func (h *APIHandler) CreateTableHandler(w http.ResponseWriter, r *http.Request) {
	var table Table
	err := json.NewDecoder(r.Body).Decode(&table)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	table.ID = uuid.New().String()
	err = h.client.SaveTable(r.Context(), &table)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(table)
}

func (h *APIHandler) GetTableByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tableID := params["id"]

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	table, err := h.client.GetTableByID(ctx, tableID)
	if err != nil {
		if errors.Is(err, ErrTableNotFound) {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(table)
}

func (h *APIHandler) BookTable(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tableID := params["id"]

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	table, err := h.client.GetTableByID(ctx, tableID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error getting table: %v", err)))
		return
	}

	if table == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Table with ID %s not found", tableID)))
		return
	}

	table.IsBooked = true
	err = h.client.SaveTable(r.Context(), table)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(table)
}
