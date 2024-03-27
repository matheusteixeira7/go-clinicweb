package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDoctor(t *testing.T) {
	doctor, err := NewDoctor("John Doe", "General Physician")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe", doctor.Name)
	assert.Equal(t, "General Physician", doctor.Specialty)
	assert.NotEmpty(t, doctor.ID)
	assert.NotZero(t, doctor.CreatedAt)
	assert.NotZero(t, doctor.UpdatedAt)

	_, err = NewDoctor("", "General Physician")
	assert.NotNil(t, err)
	assert.Equal(t, ErrDoctorNameEmpty, err)

	_, err = NewDoctor("John Doe", "")
	assert.NotNil(t, err)
	assert.Equal(t, ErrDoctorSpecialtyEmpty, err)
}
