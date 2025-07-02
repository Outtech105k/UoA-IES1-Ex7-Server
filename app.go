package main

import (
	"log"
	"net/http"

	"github.com/Outtech105k/UoA-IES1-Ex7-Server/cmd"
	"github.com/Outtech105k/UoA-IES1-Ex7-Server/routes"
	"github.com/Outtech105k/UoA-IES1-Ex7-Server/utils"
	"github.com/ory/graceful"
)

func main() {
	// Get App Config
	conf, err := cmd.ReadAppConfig("config.yaml")
	if err != nil {
		log.Fatalln("main: Failed to read config: ", err.Error())
	}
	log.Println("main: Config Read")

	appCtx := utils.AppContext{
		Config: conf,
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
