package create_doctor_usecase

import (
	"clinicweb/internal/modules/doctor/entity"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockDoctorRepository struct {
	mock.Mock
}

func (m *mockDoctorRepository) FindById(id string) (*entity.Doctor, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Doctor), args.Error(1)
}

func (m *mockDoctorRepository) Save(doctor *entity.Doctor) error {
	args := m.Called(doctor)
	return args.Error(0)
}

func TestCreateDoctorUseCase_Execute_Fail(t *testing.T) {
	mockDoctorRepository := new(mockDoctorRepository)
	mockDoctorRepository.On("Save", mock.Anything).Return(nil)

	createDoctorUseCase := NewCreateDoctorUseCase(mockDoctorRepository)

	input := CreateDoctorInputDTO{
		Name:      "",
		Specialty: "General Physician",
	}

	output, err := createDoctorUseCase.Execute(input)

	assert.NotNil(t, err)
	assert.Nil(t, output)
	assert.Equal(t, entity.ErrDoctorNameEmpty, err)

	input = CreateDoctorInputDTO{
		Name:      "John Doe",
		Specialty: "",
	}

	output, err = createDoctorUseCase.Execute(input)

	assert.NotNil(t, err)
	assert.Nil(t, output)
	assert.Equal(t, entity.ErrDoctorSpecialtyEmpty, err)
}

func TestCreateDoctorUseCase_Execute_ReturnsError(t *testing.T) {
	mockDoctorRepository := new(mockDoctorRepository)
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
	mockDoctorRepository := new(mockDoctorRepository)
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
