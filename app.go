package main

import (
	"log"
	"net/http"

	"github.com/Outtech105k/UoA-IES1-Ex7-Server/routes"
	"github.com/ory/graceful"
)

func main() {
	handler := routes.SetupRouter()
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
