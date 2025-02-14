package store

import (
	"SimpleCRUD/internal/models"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type Store struct {
	config *Config
	conn   *pgx.Conn
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	conn, err := pgx.Connect(context.Background(), s.config.DataBaseURL)
	if err != nil {
		return fmt.Errorf("ошибка при попытке подключиться на сервер: %w", err)
	}
	defer conn.Close(context.Background())

	if err := conn.Ping(context.Background()); err != nil {
		return fmt.Errorf("ошибка при попытке пингануть сервер: %w", err)
	}

	s.conn = conn

	if err := s.CreatingTable(context.Background()); err != nil {
		log.Println("ошибка при попытке создать таблицу: %w", err)
	}

	return nil
}

func (s *Store) CreatingTable(ctx context.Context) error {
	query := `
		CREATE TABLE tasks (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			description TEXT,
			status TEXT CHECK (status IN ('new', 'in_progress', 'done')) DEFAULT 'new',
			created_at TIMESTAMP DEFAULT now(),
			updated_at TIMESTAMP DEFAULT now())
	`

	_, err := s.conn.Exec(ctx, query)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) InsertTask(ctx context.Context, title string, desc string) error {
	query := `
		INSERT INTO tasks (title, description) values (@title, @description)
	`
	args := pgx.NamedArgs{
		"title":       title,
		"description": desc,
	}

	_, err := s.conn.Exec(ctx, query, args)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetTasks(ctx context.Context, tasks *[]models.Task) error {
	query := `
		SELECT * FROM tasks
	`

	rows, err := s.conn.Query(ctx, query)
	if err != nil {
		return err
	}

	log.Println(rows)
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		var id int64

		err := rows.Scan(&id, &task.Title, &task.Description, &task.Status, &task.Created_at, &task.Updated_at)
		if err != nil {
			return err
		}

		task.Id = fmt.Sprint(id)
		*tasks = append(*tasks, task)
	}

	fmt.Println(tasks)

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateTask(ctx context.Context, id int, status string) error {
	query := `
		UPDATE tasks SET status = @status
	`
	args := pgx.NamedArgs{
		"status": status,
	}
	_, err := s.conn.Exec(ctx, query, args)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteTask(ctx context.Context, id int) error {
	query := `
		DELETE FROM tasks WHERE id = @id
	`
	args := pgx.NamedArgs{
		"id": id,
	}

	_, err := s.conn.Exec(ctx, query, args)
	if err != nil {
		return err
	}

	return nil
}
