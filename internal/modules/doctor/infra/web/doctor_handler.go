package web

import (
	"clinicweb/internal/modules/doctor/infra/repository"
	"clinicweb/internal/modules/doctor/usecase/createdoctor"
	"clinicweb/internal/modules/doctor/usecase/findbyid"
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
	var dto createdoctor.CreateDoctorInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usecase := createdoctor.NewCreateDoctorUseCase(h.DoctorRepository)
	output, err := usecase.Execute(dto)
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

func (h *DoctorHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	var dto findbyid.FindDoctorByIdInputDTO
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Bad Request: 'id' parameter is required", http.StatusBadRequest)
		return
	}
	dto.ID = id
	usecase := findbyid.NewFindDoctorByIDUseCase(h.DoctorRepository)
	output, err := usecase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
