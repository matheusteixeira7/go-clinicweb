package main

import (
	"clinicweb/configs"
	"clinicweb/internal/infra/database"
	"clinicweb/internal/infra/web/webserver"
	"context"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := database.InitDatabaseClient(config.MongoDBUri, config.MongoDBName)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := db.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	done := make(chan bool)

	go func() {
		webserver.Start(config.WebServerPort, db.Database(config.MongoDBName))
		done <- true
	}()

	<-done
}
