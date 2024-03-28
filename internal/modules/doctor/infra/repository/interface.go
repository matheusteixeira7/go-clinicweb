package repository

import "clinicweb/internal/modules/doctor/entity"

type DoctorRepositoryInterface interface {
	Save(order *entity.Doctor) error
	FindById(id string) (*entity.Doctor, error)
}
