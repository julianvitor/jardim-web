package routes

import (
	"net/http"
	"jardim-web/backend/internal/handlers"
)

func SetupRoutes(mux *http.ServeMux){
	// Rotas de Admin
	mux.HandleFunc("GET /admin/devices", handlers.AdminGetDevices)
}