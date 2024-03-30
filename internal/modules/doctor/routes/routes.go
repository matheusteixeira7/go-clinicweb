package routes

import (
	"clinicweb/internal/modules/doctor/infra/repository"
	"clinicweb/internal/modules/doctor/infra/web"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func DoctorRoutes(r *chi.Mux, database *mongo.Database) {
	doctorRepository := repository.NewDoctorRepository(database.Collection("doctors"))
	handler := web.NewWebDoctorHandler(doctorRepository)

	r.Route("/doctors", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte("GET /doctors"))
			if err != nil {
				return
			}
		})

		r.Post("/", handler.CreateDoctor)
	})
}
