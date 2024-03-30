package webserver

import (
	"clinicweb/internal/modules/doctor/routes"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Start(port string, db *mongo.Database) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	routes.DoctorRoutes(r, db)

	fmt.Println("Server is running on port", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Failed to start web server: %v", err)
	}
}
