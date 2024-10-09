package main

import (
	"log"
	"yanying-api/config"
	"yanying-api/routes"
)

func main() {
	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize router
	r := routes.SetupRouter(db)

	// Start server
	if err := r.Run(":3001"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}