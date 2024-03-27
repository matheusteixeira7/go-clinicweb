package database

import (
	"clinicweb/internal/modules/doctor/entity"
	"context"

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
	coll := r.Collection.Database().Collection("doctors")
	_, err := coll.InsertOne(context.TODO(), doctor)
	if err != nil {
		return err
	}
	return nil
}
