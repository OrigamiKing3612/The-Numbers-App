package server

import "database/sql"

type Server struct {
	DB     *sql.DB
	Config *Config
}

func NewServer(DB *sql.DB, Config *Config) *Server {
	return &Server{DB, Config}
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
