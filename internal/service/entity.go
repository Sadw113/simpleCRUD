package service

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

type UpdateTaskRequest struct {
	Status string `json:"status" validate:"required"`
}
