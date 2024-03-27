package entity

import (
	"clinicweb/pkg/entity"
	"errors"
	"time"
)

type Doctor struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Specialty string    `json:"specialty"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`
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
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
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
