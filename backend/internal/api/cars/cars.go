package cars

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

func (s *Service) UpdateCheckedIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
