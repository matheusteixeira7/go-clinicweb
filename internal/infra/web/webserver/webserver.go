package webserver

import (
	"clinicweb/internal/modules/doctor/routes"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func Start(port string, db *mongo.Database) {
	mux := http.NewServeMux()

	routes.DoctorRoutes(mux, db)

	fmt.Println("Server is running on port", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Failed to start web server: %v", err)
	}
}
