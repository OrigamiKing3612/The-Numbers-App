package routes

import (
	"net/http"

	"github.com/OrigamiKing3612/The-Numbers-App/internal/api/cars"
	"github.com/OrigamiKing3612/The-Numbers-App/internal/api/users"
	"github.com/OrigamiKing3612/The-Numbers-App/internal/server"
)

func InitRoutes(server *server.Server) http.Handler {
	mux := http.NewServeMux()
	auth := users.New(server)
	cars := cars.New(server)

	mux.Handle("POST /login", auth.Login())
	mux.Handle("GET /get/cars", cars.GetAll())
	mux.Handle("PUT /update/checked_in", cars.UpdateCheckedIn())

	return mux
}
