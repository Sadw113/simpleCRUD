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

const (
	insertTaskQuery    = `INSERT INTO tasks (title, description, user_id) VALUES ($1, $2, $3) RETURNING id;`
	getTaskByIDXsQuery = `Select * FROM tasks WHERE id=$1 AND user_id = $2;`
	setStatusQuery     = `UPDATE tasks SET status = $1 WHERE id = $2 AND user_id = $3 RETURNING id;`
	deleteTaskQuery    = `DELETE FROM tasks WHERE id = $1 AND user_id = $2;`
	getTasksQuery      = `SELECT * FROM tasks WHERE user_id = $1;`
)

type repository struct {
	pool *pgxpool.Pool
}

type Repository interface {
	CreateTask(ctx context.Context, task Task) (int, error)
	GetTaskByIDXs(ctx context.Context, task Task) (*Task, error)
	UpdateTask(ctx context.Context, task Task) error
	DeleteTask(ctx context.Context, task Task) error
	GetTasks(ctx context.Context, user_id int) (map[string]Task, error)
}

func NewRepository(ctx context.Context, cfg config.PostgreSQL) (Repository, error) {
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

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse PostgreSQL config")
	}

	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeCacheDescribe

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create PostgreSQL connection pool")
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, errors.Wrap(err, "the connection doesn't ping")
	}

	return &repository{pool}, nil
}

func (r *repository) CreateTask(ctx context.Context, task Task) (int, error) {
	var id int
	err := r.pool.QueryRow(ctx, insertTaskQuery, task.Title, task.Description, task.User_id).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to insert task")
	}
	return id, nil
}

func (r *repository) GetTaskByIDXs(ctx context.Context, task Task) (*Task, error) {
	if err := r.pool.QueryRow(ctx, getTaskByIDXsQuery, task.ID, task.User_id).Scan(&task.ID, &task.User_id, &task.Title, &task.Description, &task.Status); err != nil {
		return &task, errors.New("failed to select task")
	}

	return &task, nil
}

func (r *repository) UpdateTask(ctx context.Context, task Task) error {
	var id int
	err := r.pool.QueryRow(ctx, setStatusQuery, task.Status, task.ID, task.User_id).Scan(&id)
	if err != nil {
		return errors.Wrap(err, "failed to update task")
	}

	return nil
}

func (r *repository) DeleteTask(ctx context.Context, task Task) error {
	var id int
	err := r.pool.QueryRow(ctx, deleteTaskQuery, task.ID, task.User_id).Scan(&id)
	if err != nil {
		return errors.Wrap(err, "failed to delete task")
	}

	return nil
}

func (r *repository) GetTasks(ctx context.Context, user_id int) (map[string]Task, error) {
	tasks := make(map[string]Task)
	var task Task

	rows, err := r.pool.Query(ctx, getTasksQuery, user_id)
	if err != nil {
		return tasks, errors.Wrap(err, "failed to select tasks")
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&task.ID, &task.User_id, &task.Title, &task.Description, &task.Status)
		if err != nil {
			return tasks, errors.Wrap(err, "failed to scan fields of task")
		}

		strID := strconv.Itoa(task.ID)
		tasks[strID] = task
	}

	return tasks, nil
}
