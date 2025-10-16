package users

import (
	"net/http"

	"github.com/OrigamiKing3612/The-Numbers-App/internal/server"
)

type Service struct {
	server *server.Server
}

func New(server *server.Server) *Service {
	return &Service{server}
}

func (s *Service) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Service) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Service) Remove() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Service) Reset() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
