package app

import (
	"SimpleCRUD/internal/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber"
)

func (s *APIServer) AddTask(c *fiber.Ctx) {
	task := new(models.Task)

	if err := c.BodyParser(&task); err != nil {
		c.Status(http.StatusBadRequest).Error()
		return
	}

	err := s.store.InsertTask(context.Background(), task.Title, task.Description)
	if err != nil {
		c.Status(http.StatusInternalServerError).Error()
		return
	}

	c.Status(http.StatusOK)
	c.SendString("Успех")
}

func (s *APIServer) GetTasks(c *fiber.Ctx) {
	var tasks []models.Task

	err := s.store.GetTasks(context.Background(), &tasks)
	if err != nil {
		c.Status(http.StatusInternalServerError).Error()
	}

	if tasks == nil {
		tasks = []models.Task{}
	}

	res := map[string]interface{}{
		"tasks": tasks,
	}

	c.Status(http.StatusOK).JSON(fiber.Map{"data": res})
}

func (s *APIServer) UpdateTask(c *fiber.Ctx) {
	var status string
	c.BodyParser(&status)

	if status != "in_progress" && status != "done" {
		c.Status(http.StatusBadRequest).Error()
	}

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Status(http.StatusInternalServerError).Error()
		return
	}

	err = s.store.UpdateTask(context.Background(), id, status)
	if err != nil {
		c.Status(http.StatusInternalServerError).Error()
	}

	c.Status(http.StatusOK)
}

func (s *APIServer) DeleteTask(c *fiber.Ctx) {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Status(http.StatusInternalServerError).Error()
		return
	}

	err = s.store.DeleteTask(context.Background(), id)
	if err != nil {
		c.Status(http.StatusInternalServerError).Error()
	}

	c.Status(http.StatusOK)
}
