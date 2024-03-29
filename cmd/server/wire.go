//go:build wireinject
// +build wireinject

package main

import (
	"clinicweb/internal/modules/doctor/infra/repository"
	"clinicweb/internal/modules/doctor/infra/web"
	"clinicweb/internal/modules/doctor/usecase"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var setDoctorRepositoryDependency = wire.NewSet(
	repository.NewDoctorRepository,
	wire.Bind(new(repository.DoctorRepositoryInterface), new(*repository.DoctorRepository)),
)

func NewCreateDoctorUseCase(collection *mongo.Collection) *usecase.CreateDoctorUseCase {
	wire.Build(
		setDoctorRepositoryDependency,
		usecase.NewCreateDoctorUseCase,
	)
	return &usecase.CreateDoctorUseCase{}
}

func NewWebDoctorHandler(collection *mongo.Collection) *web.DoctorHandler {
	wire.Build(
		setDoctorRepositoryDependency,
		web.NewWebDoctorHandler,
	)
	return &web.DoctorHandler{}
}
