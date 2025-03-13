package service

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

type UpdateTaskRequest struct {
	ID     int    `json:"id" validate:"required"`
	Status string `json:"status" validate:"required"`
}
