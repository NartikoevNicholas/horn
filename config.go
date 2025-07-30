package horn

import (
	cnf "github.com/NartikoevNicholas/horn/config"
)

type Config struct {
	Host       string
	Port       int
	Transport  string
	StrictMode bool
}

func DefaultConfig() Config {
	return Config{
		Host:       "127.0.0.1",
		Port:       8080,
		Transport:  cnf.TCP,
		StrictMode: true,
	}
}
