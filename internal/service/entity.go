package service

type CreateTaskRequest struct {
	User_id     int    `json:"user_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

type GetTaskByIDXsRequest struct {
	User_id int `json:"user_id" validate:"required"`
}

type UpdateTaskRequest struct {
	ID      int    `json:"id" validate:"required"`
	User_id int    `json:"user_id" validate:"required"`
	Status  string `json:"status" validate:"required"`
}

type DeleteTaskRequest struct {
	User_id int `json:"user_id" validate:"required"`
}

type GetTasksRequest struct {
	User_id int `json:"user_id" validate:"required"`
}
