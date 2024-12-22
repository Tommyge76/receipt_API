package main

import (
	"log"
	"process_receipts/internal/server"
	"process_receipts/internal/database"
)

func main() {
	db := database.CreateDatabase()
	defer db.Close()
	srv := server.NewServer(db)
	log.Fatal(srv.Start(":8080"))
} 