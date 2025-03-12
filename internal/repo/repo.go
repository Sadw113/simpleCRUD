package repo

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"

	"simple-service/internal/config"
)

// Слой репозитория, здесь должны быть все методы, связанные с базой данных

// SQL-запрос на вставку задачи
const (
	insertTaskQuery  = `INSERT INTO tasks (title, description) VALUES ($1, $2) RETURNING id;`
	getTaskByIdQuery = `Select * FROM tasks WHERE id=$1;`
	setStatusQuery   = `UPDATE tasks SET status = $1 where id = $2;`
	deleteTaskQuery  = `DELETE FROM tasks where id = $1`
	getTasksQuery    = `SELECT * FROM tasks;`
)

type repository struct {
	pool *pgxpool.Pool
}

// Repository - интерфейс с методом создания задачи
type Repository interface {
	CreateTask(ctx context.Context, task Task) (int, error)    // Создание задачи
	GetTaskByID(ctx context.Context, id uint32) (*Task, error) // Получение задачи по id
	UpdateTask(ctx context.Context, id uint32, currentStatus string) error
	DeleteTask(ctx context.Context, id uint32) error
	GetTasks(ctx context.Context) (map[string]Task, error)
}

// NewRepository - создание нового экземпляра репозитория с подключением к PostgreSQL
func NewRepository(ctx context.Context, cfg config.PostgreSQL) (Repository, error) {
	// Формируем строку подключения
	connString := fmt.Sprintf(
		`user=%s password=%s host=%s port=%d dbname=%s sslmode=%s 
        pool_max_conns=%d pool_max_conn_lifetime=%s pool_max_conn_idle_time=%s`,
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode,
		cfg.PoolMaxConns,
		cfg.PoolMaxConnLifetime.String(),
		cfg.PoolMaxConnIdleTime.String(),
	)

	// Парсим конфигурацию подключения
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse PostgreSQL config")
	}

	// Оптимизация выполнения запросов (кеширование запросов)
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeCacheDescribe

	// Создаём пул соединений с базой данных
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create PostgreSQL connection pool")
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, errors.Wrap(err, "the connection doesn't ping")
	}

	return &repository{pool}, nil
}

// CreateTask - вставка новой задачи в таблицу tasks
func (r *repository) CreateTask(ctx context.Context, task Task) (int, error) {
	var id int
	err := r.pool.QueryRow(ctx, insertTaskQuery, task.Title, task.Description).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to insert task")
	}
	return id, nil
}

func (r *repository) GetTaskByID(ctx context.Context, id uint32) (*Task, error) {
	task := Task{}

	if err := r.pool.QueryRow(ctx, getTaskByIdQuery, id).Scan(nil, &task.Title, &task.Description, &task.Status); err != nil {
		return &task, errors.New("failed to select task")
	}

	return &task, nil
}

func (r *repository) UpdateTask(ctx context.Context, id uint32, currentStatus string) error {
	_, err := r.pool.Exec(ctx, setStatusQuery, currentStatus, id)
	if err != nil {
		errors.Wrap(err, "failed to update task")
	}

	return nil
}

func (r *repository) DeleteTask(ctx context.Context, id uint32) error {
	_, err := r.pool.Exec(ctx, deleteTaskQuery, id)
	if err != nil {
		errors.Wrap(err, "failed to delete task")
	}

	return nil
}

func (r *repository) GetTasks(ctx context.Context) (map[string]Task, error) {
	tasks := make(map[string]Task)
	var task Task

	rows, err := r.pool.Query(ctx, getTasksQuery)
	if err != nil {
		errors.Wrap(err, "failed to select tasks")
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)
		if err != nil {
			errors.Wrap(err, "failed to scan fields of task")
		}

		strID := strconv.Itoa(task.ID)
		tasks[strID] = task
	}

	return tasks, nil
}
