package cars

import (
	"encoding/json"
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

func (s *Service) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cars, err := s.server.Queries.GetAllCars(r.Context())
		if err != nil {
			http.Error(w, "Failed to retrieve cars", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(cars)
	}
}
