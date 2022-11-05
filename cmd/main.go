package main

import (
	"lab4/internal/handler"
	"lab4/internal/service"
	"lab4/internal/storage"
	"log"
)

func main() {
	db, err := storage.CreateDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("can't close db err: %v\n", err)
		} else {
			log.Printf("db closed")
		}
	}()
	storage := storage.NewStorage(db)
	service := service.NewService(storage)
	handler := handler.NewHandler(service)

	handler.InitRoutes()
}
