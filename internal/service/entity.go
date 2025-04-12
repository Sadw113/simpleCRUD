package service

type CreateTaskRequest struct {
	UserID      int    `json:"user_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

type GetTaskByIDXsRequest struct {
	UserID int `json:"user_id" validate:"required"`
}

type UpdateTaskRequest struct {
	ID     string `json:"id" validate:"required"`
	UserID int    `json:"user_id" validate:"required"`
	Status string `json:"status" validate:"required"`
}

type DeleteTaskRequest struct {
	UserID int `json:"user_id" validate:"required"`
}

type GetTasksRequest struct {
	UserID int `json:"user_id" validate:"required"`
}
