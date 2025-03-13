package service

import (
	"encoding/json"
	"simple-service/internal/dto"
	"simple-service/internal/repo"
	"simple-service/pkg/validator"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Слой бизнес-логики. Тут должна быть основная логика сервиса

// Service - интерфейс для бизнес-логики
type Service interface {
	CreateTask(ctx *fiber.Ctx) error
	GetTaskByID(ctx *fiber.Ctx) error
	DeleteTask(ctx *fiber.Ctx) error
	UpdateTask(ctx *fiber.Ctx) error
	GetTasks(ctx *fiber.Ctx) error
}

type service struct {
	repo repo.Repository
	log  *zap.SugaredLogger
}

// NewService - конструктор сервиса
func NewService(repo repo.Repository, logger *zap.SugaredLogger) Service {
	return &service{
		repo: repo,
		log:  logger,
	}
}

// CreateTask - обработчик запроса на создание задачи
func (s *service) CreateTask(ctx *fiber.Ctx) error {
	var req CreateTaskRequest

	if err := json.Unmarshal(ctx.Body(), &req); err != nil {
		s.log.Error("Invalid request body", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Invalid request body")
	}

	if vErr := validator.Validate(ctx.Context(), req); vErr != nil {
		return dto.BadResponseError(ctx, dto.FieldIncorrect, vErr.Error())
	}

	task := repo.Task{
		Title:       req.Title,
		Description: req.Description,
	}
	taskID, err := s.repo.CreateTask(ctx.Context(), task)
	if err != nil {
		s.log.Error("Failed to insert task", zap.Error(err))
		return dto.InternalServerError(ctx)
	}

	response := dto.StatusOK(ctx, map[string]int{"task_id": taskID})

	return response
}

func (s *service) GetTaskByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		s.log.Error("Failed to parse id from request", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Invalid request body")
	}

	task, err := s.repo.GetTaskByID(ctx.Context(), uint32(id))
	if err != nil {
		s.log.Error("Failed to selected from DB Table", zap.Error(err))
		return dto.InternalServerError(ctx)
	}

	response := dto.StatusOK(ctx, task)

	return response
}

func (s *service) UpdateTask(ctx *fiber.Ctx) error {
	var req UpdateTaskRequest

	if err := json.Unmarshal(ctx.Body(), &req); err != nil {
		s.log.Error("Invalid request body", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Invalid request body")
	}

	if vErr := validator.Validate(ctx.Context(), req); vErr != nil {
		return dto.BadResponseError(ctx, dto.FieldIncorrect, vErr.Error())
	}

	task := repo.Task{
		ID:     req.ID,
		Status: req.Status,
	}

	err := s.repo.UpdateTask(ctx.Context(), task)
	if err != nil {
		s.log.Error("Failed to update table", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Failed to update table")
	}

	response := dto.StatusOK(ctx, "Success updating")

	return response
}

func (s *service) DeleteTask(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		s.log.Error("Failed to parse id from request", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Invalid request body")
	}

	err = s.repo.DeleteTask(ctx.Context(), uint32(id))
	if err != nil {
		s.log.Error("Failed to delete from table", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Failed to delete from table")
	}

	response := dto.StatusOK(ctx, "Delete was successful")

	return response
}

func (s *service) GetTasks(ctx *fiber.Ctx) error {
	tasks, err := s.repo.GetTasks(ctx.Context())
	if err != nil {
		s.log.Error("Failed to select all tasks from table", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Failed to select all tasks from table")
	}

	response := dto.StatusOK(ctx, tasks)

	return response
}
