package server

import (
	"database/sql"

	"github.com/OrigamiKing3612/The-Numbers-App/internal/database"
)

type Server struct {
	DB      *sql.DB
	Config  *Config
	Queries *database.Queries
}

func NewServer(DB *sql.DB, Config *Config, Queries *database.Queries) *Server {
	return &Server{DB, Config, Queries}
}

type Config struct {
	Debug bool
}

func NewConfig(Debug bool) *Config {
	return &Config{
		Debug,
	}
}

func (s *Server) IsDebug() bool {
	return s.Config.Debug
}
