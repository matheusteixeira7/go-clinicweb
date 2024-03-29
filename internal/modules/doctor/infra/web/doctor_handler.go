package web

import (
	"clinicweb/internal/modules/doctor/infra/repository"
	"clinicweb/internal/modules/doctor/usecase"
	"encoding/json"
	"net/http"
)

type DoctorHandler struct {
	DoctorRepository repository.DoctorRepositoryInterface
}

func NewWebDoctorHandler(doctorRepository repository.DoctorRepositoryInterface) *DoctorHandler {
	return &DoctorHandler{
		DoctorRepository: doctorRepository,
	}
}

func (h *DoctorHandler) CreateDoctor(w http.ResponseWriter, r *http.Request) {
	var dto usecase.CreateDoctorInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createDoctor := usecase.NewCreateDoctorUseCase(h.DoctorRepository)
	output, err := createDoctor.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
