package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"simple-service/internal/api/middleware"
	"simple-service/internal/service"
)

type Routers struct {
	Service service.Service
}

func NewRouters(r *Routers, token string) *fiber.App {
	app := fiber.New()

	// Настройка CORS (разрешенные методы, заголовки, авторизация)
	app.Use(cors.New(cors.Config{
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowHeaders:     "Accept, Authorization, Content-Type, X-CSRF-Token, X-REQUEST-SomeID",
		ExposeHeaders:    "Link",
		AllowCredentials: false,
		MaxAge:           300,
	}))

	apiGroupTask := app.Group("/task", middleware.Authorization(token))
	apiGroupTask.Post("/create_task", r.Service.CreateTask)
	apiGroupTask.Get("/get_task/:id", r.Service.GetTaskByID)
	apiGroupTask.Put("/update_task", r.Service.UpdateTask)
	apiGroupTask.Delete("/delete_task/:id", r.Service.DeleteTask)

	apiGroupTasks := app.Group("/tasks", middleware.Authorization(token))
	apiGroupTasks.Get("/get_tasks", r.Service.GetTasks)

	return app
}
