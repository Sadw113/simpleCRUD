package app

import (
	"SimpleCRUD/store"

	"github.com/gofiber/fiber"
)

type APIServer struct {
	config *Config
	store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

func (s *APIServer) Start() error {
	router := fiber.New()
	s.configurateRouter(router)

	if err := s.configurateStore(); err != nil {
		return err
	}

	return router.Listen(s.config.BindAddr)
}

func (s *APIServer) configurateRouter(router fiber.Router) {
	note := router.Group("/tasks")

	note.Post("", s.AddTask)
	note.Get("", s.GetTasks)
	note.Put(":id", s.UpdateTask)
	note.Delete(":id", s.DeleteTask)
}

func (s *APIServer) configurateStore() error {
	store := store.New(s.config.Store)

	if err := store.Open(); err != nil {
		return err
	}
	s.store = store

	return nil
}
