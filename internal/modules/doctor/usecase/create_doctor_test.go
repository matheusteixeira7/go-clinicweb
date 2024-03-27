package usecase

import (
	doctorEntity "clinicweb/internal/modules/doctor/entity"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDoctorRepository struct {
	mock.Mock
}

func (m *MockDoctorRepository) Save(doctor *doctorEntity.Doctor) error {
	args := m.Called(doctor)
	return args.Error(0)
}

func TestCreateDoctorUseCase_Execute_Fail(t *testing.T) {
	mockDoctorRepository := new(MockDoctorRepository)
	mockDoctorRepository.On("Save", mock.Anything).Return(nil)

	createDoctorUseCase := NewCreateDoctorUseCase(mockDoctorRepository)

	input := CreateDoctorInputDTO{
		Name:      "",
		Specialty: "General Physician",
	}

	output, err := createDoctorUseCase.Execute(input)

	assert.NotNil(t, err)
	assert.Nil(t, output)
	assert.Equal(t, doctorEntity.ErrDoctorNameEmpty, err)

	input = CreateDoctorInputDTO{
		Name:      "John Doe",
		Specialty: "",
	}

	output, err = createDoctorUseCase.Execute(input)

	assert.NotNil(t, err)
	assert.Nil(t, output)
	assert.Equal(t, doctorEntity.ErrDoctorSpecialtyEmpty, err)
}

func TestCreateDoctorUseCase_Execute_ReturnsError(t *testing.T) {
	mockDoctorRepository := new(MockDoctorRepository)
	mockDoctorRepository.On("Save", mock.Anything).Return(errors.New("error"))

	createDoctorUseCase := NewCreateDoctorUseCase(mockDoctorRepository)

	input := CreateDoctorInputDTO{
		Name:      "John Doe",
		Specialty: "General Physician",
	}

	output, err := createDoctorUseCase.Execute(input)

	assert.NotNil(t, err)
	assert.Nil(t, output)
	assert.EqualError(t, err, "error")
}

func TestCreateDoctorUseCase_Execute_Success(t *testing.T) {
	mockDoctorRepository := new(MockDoctorRepository)
	mockDoctorRepository.On("Save", mock.Anything).Return(nil)

	createDoctorUseCase := NewCreateDoctorUseCase(mockDoctorRepository)

	input := CreateDoctorInputDTO{
		Name:      "John Doe",
		Specialty: "General Physician",
	}

	output, err := createDoctorUseCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, input.Name, output.Name)
	assert.Equal(t, input.Specialty, output.Specialty)
	assert.NotEmpty(t, output.ID)
	assert.NotZero(t, output.CreatedAt)
	assert.NotZero(t, output.UpdatedAt)
}
