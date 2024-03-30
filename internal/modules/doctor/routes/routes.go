package routes

import (
	"clinicweb/internal/modules/doctor/infra/repository"
	"clinicweb/internal/modules/doctor/infra/web"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func DoctorRoutes(mux *http.ServeMux, database *mongo.Database) {
	doctorRepository := repository.NewDoctorRepository(database.Collection("doctors"))
	handler := web.NewWebDoctorHandler(doctorRepository)

	mux.HandleFunc("GET /doctors/{id}", handler.FindByID)
	mux.HandleFunc("POST /doctors", handler.CreateDoctor)
}
