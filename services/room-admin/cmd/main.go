package main

import "musical-tables-api/services/room-admin/internal/api"

func main() {
	server := api.NewServer()
	server.ConfigureEndpoints()
	server.Start()
}
