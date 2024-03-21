package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lipzpoom/controller"
)

func NewRouter(noteController *controller.NoteController) *fiber.App {
	router := fiber.New()

	router.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "welcome to Golang, fiber, and GOrm",
		})
	})

	router.Route("/notes", func(router fiber.Router) {
		router.Post("/", noteController.Create)
		router.Get("", noteController.FindAll)
	})

	router.Route("/notes/:noteId", func(router fiber.Router) {
		router.Delete("", noteController.Delete)
		router.Get("", noteController.FindById)
		router.Patch("", noteController.Update)
	})

	return router
}
