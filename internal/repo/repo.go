package repo

import (
	"context"

	"github.com/pkg/errors"
)

// Слой репозитория, здесь должны быть все методы, связанные с базой данных

type repository struct {
	tasks map[int]Task
}

// Repository - интерфейс с методом создания задачи
type Repository interface {
	CreateTask(ctx context.Context, task Task) (int, error)            // Создание задачи
	GetTasks(ctx context.Context) map[int]Task                         // Получение задач
	GetTaskByID(ctx context.Context, id uint32) (*Task, error)         // Получение задачи по id
	UpdateTask(ctx context.Context, id uint32, task Task) (int, error) // Обновление задачи
	DeleteTask(ctx context.Context, id uint32) (int, error)            // Удаление задачи
}

// NewRepository - создание нового экземпляра репозитория с подключением к PostgreSQL
func NewRepository(ctx context.Context) (Repository, error) {
	// создаем мапу для хранения данных в памяти
	var tasks = make(map[int]Task)

	return &repository{tasks: tasks}, nil
}

// CreateTask - вставка новой задачи в таблицу tasks
func (r *repository) CreateTask(ctx context.Context, task Task) (int, error) {
	id := task.ID

	_, exist := r.tasks[id]

	if exist {
		return 0, errors.New("This task is exist")
	}

	r.tasks[id] = task

	return id, nil
}

func (r *repository) GetTaskByID(ctx context.Context, id uint32) (*Task, error) {
	task, exist := r.tasks[int(id)]

	if !exist {
		return nil, errors.New("This task is not exist")
	}

	return &task, nil
}

func (r *repository) GetTasks(ctx context.Context) map[int]Task {
	return r.tasks
}

func (r *repository) UpdateTask(ctx context.Context, id uint32, task Task) (int, error) {
	_, exist := r.tasks[int(id)]

	if !exist {
		return 0, errors.New("This task is not exist")
	}

	r.tasks[int(id)] = task

	return int(id), nil
}

func (r *repository) DeleteTask(ctx context.Context, id uint32) (int, error) {
	_, exist := r.tasks[int(id)]

	if !exist {
		return 0, errors.New("This task is not exist")
	}

	delete(r.tasks, int(id))

	return int(id), nil
}
