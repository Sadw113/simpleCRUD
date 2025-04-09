package service

// import (
// 	"bytes"
// 	"encoding/json"
// 	"errors"
// 	"net/http"
// 	"testing"

// 	"simple-service/internal/dto"
// 	"simple-service/internal/repo"
// 	"simple-service/internal/repo/mocks"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"go.uber.org/zap"
// )

// func TestCreateTask(t *testing.T) {
// 	mockRepo := new(mocks.Repository)
// 	logger := zap.NewNop().Sugar()

// 	s := NewService(mockRepo, logger)

// 	app := fiber.New()
// 	app.Post("/tasks", s.CreateTask)

// 	t.Run("успешное создание задачи", func(t *testing.T) {
// 		task := CreateTaskRequest{
// 			Title:       "Test Task",
// 			Description: "Test Description",
// 		}
// 		body, _ := json.Marshal(task)

// 		mockRepo.On("CreateTask", mock.Anything, repo.Task{
// 			Title:       task.Title,
// 			Description: task.Description,
// 		}).Return(1, nil).Once()

// 		req, err := http.NewRequest("POST", "/tasks", bytes.NewReader(body))
// 		assert.NoError(t, err)
// 		req.Header.Set("Content-Type", "application/json")

// 		resp, err := app.Test(req)
// 		assert.NoError(t, err)
// 		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

// 		var response dto.Response
// 		json.NewDecoder(resp.Body).Decode(&response)
// 		assert.Equal(t, "success", response.Status)

// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("ошибка валидации входных данных", func(t *testing.T) {
// 		body := []byte(`{}`)

// 		req, err := http.NewRequest("POST", "/tasks", bytes.NewReader(body))
// 		assert.NoError(t, err)
// 		req.Header.Set("Content-Type", "application/json")

// 		resp, err := app.Test(req)
// 		assert.NoError(t, err)
// 		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

// 		var response dto.Response
// 		json.NewDecoder(resp.Body).Decode(&response)
// 		assert.Equal(t, "error", response.Status)
// 	})

// 	t.Run("ошибка при создании задачи в БД", func(t *testing.T) {
// 		task := CreateTaskRequest{
// 			Title:       "Test Task",
// 			Description: "Test Description",
// 		}
// 		body, _ := json.Marshal(task)

// 		mockRepo.On("CreateTask", mock.Anything, repo.Task{
// 			Title:       task.Title,
// 			Description: task.Description,
// 		}).Return(0, errors.New("DB error")).Once()

// 		req, _ := http.NewRequest("POST", "/tasks", bytes.NewReader(body))
// 		req.Header.Set("Content-Type", "application/json")

// 		resp, err := app.Test(req)
// 		assert.NoError(t, err)
// 		assert.NoError(t, err)
// 		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

// 		var response dto.Response
// 		json.NewDecoder(resp.Body).Decode(&response)
// 		assert.Equal(t, "error", response.Status)

// 		mockRepo.AssertExpectations(t)
// 	})
// }

// func TestGetTaskByID(t *testing.T) {
// 	mockRepo := new(mocks.Repository)
// 	logger := zap.NewNop().Sugar()

// 	s := NewService(mockRepo, logger)

// 	app := fiber.New()
// 	app.Get("/get_task/:id", s.GetTaskByID)

// 	t.Run("ожидаем ошибку парсинга id", func(t *testing.T) {
// 	})
// }
