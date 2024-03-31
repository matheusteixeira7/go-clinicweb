package createdoctor

import (
	"clinicweb/internal/modules/doctor/entity"
	"clinicweb/internal/modules/doctor/infra/repository"
)

type CreateDoctorInputDTO struct {
	Name      string `json:"name"`
	Specialty string `json:"specialty"`
}

type CreateDoctorOutputDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Specialty string `json:"specialty"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type CreateDoctorUseCase struct {
	repository repository.DoctorRepositoryInterface
}

func NewCreateDoctorUseCase(repository repository.DoctorRepositoryInterface) *CreateDoctorUseCase {
	return &CreateDoctorUseCase{
		repository: repository,
	}
}
func (c *CreateDoctorUseCase) Execute(input CreateDoctorInputDTO) (*CreateDoctorOutputDTO, error) {
	doctor, err := entity.NewDoctor(input.Name, input.Specialty)
	if err != nil {
		return nil, err
	}
	err = c.repository.Save(doctor)
	if err != nil {
		return nil, err
	}
	return &CreateDoctorOutputDTO{
		ID:        doctor.ID,
		Name:      doctor.Name,
		Specialty: doctor.Specialty,
		CreatedAt: doctor.CreatedAt,
		UpdatedAt: doctor.UpdatedAt,
	}, nil
}
