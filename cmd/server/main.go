package main

import (
	"clinicweb/configs"
	"clinicweb/internal/infra/web/webserver"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	webserver := webserver.NewWebServer(configs.WebServerPort)
	webserver.AddHandler("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	fmt.Println("Starting web server on port", configs.WebServerPort)
	webserver.Start()
}
