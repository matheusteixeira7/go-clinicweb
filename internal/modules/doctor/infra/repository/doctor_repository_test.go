//go:build integration

package repository

import (
	"clinicweb/internal/modules/doctor/entity"
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	resource, err := pool.Run("mongo", "latest", []string{"MONGO_INITDB_ROOT_USERNAME=root", "MONGO_INITDB_ROOT_PASSWORD=secret"})
	if err != nil {
		pool.Purge(resource)
		log.Fatalf("Could not start resource: %s", err)
	}

	if err := pool.Retry(func() error {
		var err error
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		db, err = mongo.Connect(ctx, options.Client().ApplyURI(
			fmt.Sprintf("mongodb://root:secret@localhost:%s", resource.GetPort("27017/tcp")),
		))

		if err != nil {
			return err
		}
		return db.Ping(ctx, nil)
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestSaveDoctor_Fail(t *testing.T) {
	doctor, _ := entity.NewDoctor("", "General Physician")
	doctorRepository := NewDoctorRepository(db.Database("clinicweb").Collection("doctors"))
	err := doctorRepository.Save(doctor)
	assert.NotNil(t, err)
	assert.Error(t, err)

	doctor, _ = entity.NewDoctor("John Doe", "")
	err = doctorRepository.Save(doctor)
	assert.NotNil(t, err)
	assert.Error(t, err)
}

func TestSaveDoctor_Success(t *testing.T) {
	doctor, _ := entity.NewDoctor("John Doe", "General Physician")
	doctorRepository := NewDoctorRepository(db.Database("clinicweb").Collection("doctors"))
	err := doctorRepository.Save(doctor)
	assert.Nil(t, err)

	foundDoctor, err := doctorRepository.FindById(doctor.ID)
	assert.Nil(t, err)
	assert.Equal(t, doctor.ID, foundDoctor.ID)
	assert.Equal(t, doctor.Name, foundDoctor.Name)
	assert.Equal(t, doctor.Specialty, foundDoctor.Specialty)
	assert.Equal(t, doctor.CreatedAt, foundDoctor.CreatedAt)
	assert.Equal(t, doctor.UpdatedAt, foundDoctor.UpdatedAt)
}

func TestFindDoctor_Fail(t *testing.T) {
	doctorRepository := NewDoctorRepository(db.Database("clinicweb").Collection("doctors"))

	foundDoctor, err := doctorRepository.FindById("invalid_id")
	assert.NotNil(t, err)
	assert.Nil(t, foundDoctor)
	assert.EqualError(t, err, "mongo: no documents in result")
}

func TestFindDoctor_Success(t *testing.T) {
	doctor, _ := entity.NewDoctor("John Doe", "General Physician")
	doctorRepository := NewDoctorRepository(db.Database("clinicweb").Collection("doctors"))
	err := doctorRepository.Save(doctor)
	assert.Nil(t, err)

	foundDoctor, err := doctorRepository.FindById(doctor.ID)
	assert.Nil(t, err)
	assert.Equal(t, doctor.ID, foundDoctor.ID)
	assert.Equal(t, doctor.Name, foundDoctor.Name)
	assert.Equal(t, doctor.Specialty, foundDoctor.Specialty)
	assert.Equal(t, doctor.CreatedAt, foundDoctor.CreatedAt)
	assert.Equal(t, doctor.UpdatedAt, foundDoctor.UpdatedAt)
}
