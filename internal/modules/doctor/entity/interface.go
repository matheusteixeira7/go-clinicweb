package entity

type DoctorRepositoryInterface interface {
	Save(order *Doctor) error
}
