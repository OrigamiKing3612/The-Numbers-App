package routes

import (
	"net/http"

	"github.com/OrigamiKing3612/The-Numbers-App/internal/server"
)

func InitRoutes(server *server.Server) http.Handler {
	mux := http.NewServeMux()

	return mux
}
