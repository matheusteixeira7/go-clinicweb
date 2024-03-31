package findbyid

import "clinicweb/internal/modules/doctor/infra/repository"

type FindDoctorByIdInputDTO struct {
	ID string `json:"id"`
}

type FindDoctorByIdOutputDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Specialty string `json:"specialty"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type FindDoctorByIDUseCase struct {
	repository repository.DoctorRepositoryInterface
}

func NewFindDoctorByIDUseCase(repository repository.DoctorRepositoryInterface) *FindDoctorByIDUseCase {
	return &FindDoctorByIDUseCase{
		repository: repository,
	}
}

func (f *FindDoctorByIDUseCase) Execute(input FindDoctorByIdInputDTO) (*FindDoctorByIdOutputDTO, error) {
	doctor, err := f.repository.FindById(input.ID)
	if err != nil {
		return nil, err
	}
	return &FindDoctorByIdOutputDTO{
		ID:        doctor.ID,
		Name:      doctor.Name,
		Specialty: doctor.Specialty,
		CreatedAt: doctor.CreatedAt,
		UpdatedAt: doctor.UpdatedAt,
	}, nil
}
