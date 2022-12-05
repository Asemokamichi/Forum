package main

import (
	"log"

	"github.com/Asemokamichi/Forum/internal/delivery"
	"github.com/Asemokamichi/Forum/internal/repository"
	"github.com/Asemokamichi/Forum/internal/server"
	"github.com/Asemokamichi/Forum/internal/service"
)

func main() {
	db, err := repository.NewDataBase("db.db")
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewRepository(db)
	if err := repository.Create(); err != nil {
		log.Fatal(err)
		return
	}

	service := service.NewService(repository)

	handler := delivery.NewHandler(service)

	server := server.NewServer(handler)

	log.Printf("Starting server at port \nhttp://localhost:8080\n")
	if err := server.S.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
