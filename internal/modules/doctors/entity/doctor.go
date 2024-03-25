package entity

import (
	"clinicweb/pkg/entity"
	"errors"
)

type Doctor struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Specialty string    `json:"specialty"`
}

var (
	ErrDoctorNameEmpty      = errors.New("doctor name can't be empty")
	ErrDoctorSpecialtyEmpty = errors.New("doctor specialty can't be empty")
)

func NewDoctor(name, specialty string) (*Doctor, error) {
	doctor := &Doctor{
		ID:        entity.NewID(),
		Name:      name,
		Specialty: specialty,
	}
	err := doctor.Validate()
	if err != nil {
		return nil, err
	}
	return doctor, nil
}

func (d *Doctor) Validate() error {
	if d.Name == "" {
		return ErrDoctorNameEmpty
	}
	if d.Specialty == "" {
		return ErrDoctorSpecialtyEmpty
	}
	return nil
}
