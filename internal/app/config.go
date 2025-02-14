package app

import "SimpleCRUD/store"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		Store:    store.NewConfig(),
	}
}
