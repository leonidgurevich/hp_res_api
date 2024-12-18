package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/leonidgurevich/hc_rest_api/config"
	"github.com/leonidgurevich/hc_rest_api/routers"

)

func main() {
	// Load configuration

	configPath := "/home/leonid/hc_rest_api/config/connect.json"
	config, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Unable to load configuration: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	fmt.Printf("Database DSN: %s\n", dsn)
	// Initialize the database
	db.Connect(dsn)

	// Set up the router
	router := controllers.SetupRouter()

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
