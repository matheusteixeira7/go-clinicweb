package main

import (
	"clinicweb/configs"
	"clinicweb/internal/infra/database"
	"clinicweb/internal/infra/web/webserver"
	"fmt"
	"log"
	"net/http"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	_, err = database.InitDatabaseClient(configs.MongoDBUri)
	if err != nil {
		panic(err)
	}
	webServerPort := configs.WebServerPort
	if webServerPort == "" {
		log.Fatal("You must set your 'WEB_SERVER_PORT' environment variable.")
	}
	webserver := webserver.NewWebServer(webServerPort)
	webserver.AddHandler("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	fmt.Println("Starting web server on port", webServerPort)
	webserver.Start()
}
