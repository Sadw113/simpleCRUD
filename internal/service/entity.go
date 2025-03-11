package service

// TaskRequest - структура, представляющая тело запроса
type TaskRequest struct {
	Id          int    `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Status      string `json:"status" validate:"required"`
}
