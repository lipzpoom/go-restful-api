package main

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/lipzpoom/config"
	"github.com/lipzpoom/controller"
	"github.com/lipzpoom/model"
	"github.com/lipzpoom/repository"
	"github.com/lipzpoom/router"
	"github.com/lipzpoom/service"
)

func main() {
	fmt.Println("Run Service ...")

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load environment variables\n", err)
	}

	// Database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("notes").AutoMigrate(&model.Note{})

	// Init Repository
	noteRepository := repository.NewNoteRepositoryImpl(db)

	// Init Service
	noteService := service.NewNoteServiceImpl(noteRepository, validate)

	// Init Controller
	noteController := controller.NewNoteController(noteService)

	// Routes
	routes := router.NewRouter(noteController)

	app := fiber.New()

	app.Mount("/api", routes)
	log.Fatal(app.Listen(":8000"))
}

// router := fiber.New()
// app := fiber.New()

// app.Mount("/api",router)

// router.Get("/healthchecker",func(c *fiber.Ctx) error{
// 	return c.Status(200).JSON(fiber.Map{
// 		"status":"success",
// 		"message":"welcome to Golang, fiber, and GOrm",
// 	})
// })
