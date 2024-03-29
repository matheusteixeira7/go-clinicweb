package main

import (
	"clinicweb/configs"
	"clinicweb/internal/infra/web/webserver"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoDBUri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	var result bson.M
	if err := client.Database(config.MongoDBName).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	dbDoctorsCollection := client.Database(config.MongoDBName).Collection("doctors")
	NewCreateDoctorUseCase(dbDoctorsCollection)
	webDoctorHandler := NewWebDoctorHandler(dbDoctorsCollection)
	web := webserver.NewWebServer(config.WebServerPort)
	web.AddHandler("/doctors", webDoctorHandler.CreateDoctor)
	fmt.Println("Starting web server on port", config.WebServerPort)
	web.Start()
}
