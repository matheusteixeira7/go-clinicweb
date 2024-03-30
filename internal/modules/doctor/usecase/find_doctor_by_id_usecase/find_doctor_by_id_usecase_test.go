package find_doctor_by_id_usecase

import (
	"clinicweb/internal/modules/doctor/entity"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockDoctorRepository struct {
	mock.Mock
}

func (m *mockDoctorRepository) FindById(id string) (*entity.Doctor, error) {
	args := m.Called(id)
	var doc *entity.Doctor
	if args.Get(0) != nil {
		doc = args.Get(0).(*entity.Doctor)
	}
	return doc, args.Error(1)
}

func (m *mockDoctorRepository) Save(doctor *entity.Doctor) error {
	args := m.Called(doctor)
	return args.Error(0)
}

func TestFindDoctorByIdUseCase_Execute_Fail(t *testing.T) {
	mockDoctorRepository := new(mockDoctorRepository)
	mockDoctorRepository.On("FindById", mock.Anything).Return(nil, errors.New("error"))

	findDoctorByIDUseCase := NewFindDoctorByIDUseCase(mockDoctorRepository)

	input := FindDoctorByIdInputDTO{
		ID: "uuid",
	}

	output, err := findDoctorByIDUseCase.Execute(input)

	assert.NotNil(t, err)
	assert.Nil(t, output)
	assert.EqualError(t, err, "error")
}

func TestFindDoctorByIDUseCase_Execute_Success(t *testing.T) {
	mockDoctor := entity.Doctor{
		Name:      "John Doe",
		Specialty: "Specialty",
	}

	mockDoctorRepository := new(mockDoctorRepository)
	mockDoctorRepository.On("FindById", mockDoctor.ID).Return(&mockDoctor, nil)

	findDoctorByIDUseCase := NewFindDoctorByIDUseCase(mockDoctorRepository)

	output, err := findDoctorByIDUseCase.Execute(FindDoctorByIdInputDTO{
		ID: mockDoctor.ID,
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, mockDoctor.ID, output.ID)
	assert.Equal(t, mockDoctor.Name, output.Name)
	assert.Equal(t, mockDoctor.Specialty, output.Specialty)
	assert.Equal(t, mockDoctor.CreatedAt, output.CreatedAt)
	assert.Equal(t, mockDoctor.UpdatedAt, output.UpdatedAt)
	mockDoctorRepository.AssertCalled(t, "FindById", mockDoctor.ID)
}
