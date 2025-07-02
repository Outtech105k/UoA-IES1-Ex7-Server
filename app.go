package main

import (
	"log"
	"net/http"

	"github.com/Outtech105k/UoA-IES1-Ex7-Server/cmd"
	"github.com/Outtech105k/UoA-IES1-Ex7-Server/routes"
	"github.com/Outtech105k/UoA-IES1-Ex7-Server/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ory/graceful"
)

func main() {
	// Get App Config
	conf, err := cmd.ReadAppConfig("config.yaml")
	if err != nil {
		log.Fatalln("main: Failed to read config: ", err.Error())
	}
	log.Println("main: Config Read")

	// Connect DB
	db, err := sqlx.Open("sqlite3", conf.Sqlite3.Path)
	if err != nil {
		log.Fatalln("main: Failed to open SQLite3: ", err.Error())
	}
	defer db.Close()

	// DB initialize
	if _, err := db.Exec(`
CREATE TABLE IF NOT EXISTS histories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    status TEXT NOT NULL,
    detected DATETIME NOT NULL
);
`); err != nil {
		log.Panicln("main: Failed to initialize DB: ", err.Error())
	}

	appCtx := utils.AppContext{
		Config: conf,
		DB:     db,
	}

	// Setup and Run Server
	handler := routes.SetupRouter(appCtx)
	server := graceful.WithDefaults(&http.Server{
		Addr:    ":8080",
		Handler: handler,
	})

	log.Println("main: Starting the server")
	if err := graceful.Graceful(server.ListenAndServe, server.Shutdown); err != nil {
		log.Fatalln("main: Failed to gracefully shutdown: ", err.Error())
	}
	log.Println("main: Server was shutdown gracefully")
}
