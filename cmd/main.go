package main

import (
	"github.com/MaximGribkov/toDO/pkg/handler"
	"log"

	project "github.com/MaximGribkov/toDO"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(project.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error start server http: %s", err.Error())
	}
}
