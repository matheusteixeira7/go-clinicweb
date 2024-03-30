package web

import (
	"clinicweb/internal/modules/doctor/infra/repository"
	"clinicweb/internal/modules/doctor/usecase/create_doctor_usecase"
	"clinicweb/internal/modules/doctor/usecase/find_doctor_by_id_usecase"
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
	var dto create_doctor_usecase.CreateDoctorInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usecase := create_doctor_usecase.NewCreateDoctorUseCase(h.DoctorRepository)
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
	var dto find_doctor_by_id_usecase.FindDoctorByIdInputDTO
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Bad Request: 'id' parameter is required", http.StatusBadRequest)
		return
	}
	dto.ID = id
	usecase := find_doctor_by_id_usecase.NewFindDoctorByIDUseCase(h.DoctorRepository)
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
