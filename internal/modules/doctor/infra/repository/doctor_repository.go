package repository

import (
	"clinicweb/internal/modules/doctor/entity"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DoctorRepository struct {
	Collection *mongo.Collection
}

func NewDoctorRepository(collection *mongo.Collection) *DoctorRepository {
	return &DoctorRepository{
		Collection: collection,
	}
}

func (r *DoctorRepository) Save(doctor *entity.Doctor) error {
	_, err := r.Collection.InsertOne(context.TODO(), &doctor)
	if err != nil {
		return err
	}
	return nil
}

func (r *DoctorRepository) FindById(id string) (*entity.Doctor, error) {
	var doctor entity.Doctor
	err := r.Collection.FindOne(context.TODO(), bson.D{{Key: "id", Value: id}}).Decode(&doctor)
	if err != nil {
		return nil, err
	}
	return &doctor, nil
}
