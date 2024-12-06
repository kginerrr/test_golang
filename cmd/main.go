package main

import (
	"fibertesttask/internal/db"
	"fibertesttask/internal/handler"
	"fibertesttask/internal/repository"
	"fibertesttask/internal/service"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var onceDB sync.Once

func main() {

	db.InitDatabase()

	repo := repository.NewUserRepository(db.DB)
	serv := service.NewUserService(repo)
	h := handler.NewUserHandler(serv)

	app := fiber.New()

	app.Post("/users", h.Create)
	app.Get("/users", h.GetAll)
	app.Get("/users/:id", h.GetByID)
	app.Put("/users/:id", h.Update)
	app.Delete("/users/:id", h.Delete)

	err := app.Listen("127.0.0.1:8080")
	if err != nil {
		return
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	log.Println("Shutting down server...")
	db.Close()
	log.Println("db closed")
	log.Println("server stopped")

}
