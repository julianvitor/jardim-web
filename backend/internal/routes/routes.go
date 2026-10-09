package routes

import (
	"net/http"
	"jardim-web/backend/internal/handlers"
)

func SetupRoutes(mux *http.ServeMux){
	// Rotas de Admin
	mux.HandleFunc("GET /admin/devices", handlers.AdminGetDevices)
	mux.HandleFunc("POST /admin/devices", handlers.AdminPostDevices)
	mux.HandleFunc("PATCH /admin/devices", handlers.AdminPatchDevices)
	mux.HandleFunc("DELETE /admin/devices", handlers.AdminDeleteDevices)
	mux.HandleFunc("GET /admin/users", handlers.AdminGetUsers)
	mux.HandleFunc("DELETE /admin/users", handlers.AdminDeleteUsers)
	mux.HandleFunc("GET /admin/telemetry", handlers.AdminGetTelemetry)
	mux.HandleFunc("DELETE /admin/telemetry", handlers.AdminDeleteTelemetry)
	// Rotas de Hardware
	mux.HandleFunc("POST /hardware/telemetry", handlers.HardwarePostTelemetry)
	mux.HandleFunc("GET /hardware/settings", handlers.HardwareGetSettings)
	// Rotas de Users
	mux.HandleFunc("GET /users/devices", handlers.UsersGetDevices)
	mux.HandleFunc("POST /users/devices", handlers.UsersPostDevices)
	mux.HandleFunc("PATCH /users/devices", handlers.UsersPatchDevices)
	mux.HandleFunc("DELETE /users/devices", handlers.UsersDeleteDevices)
	mux.HandleFunc("GET /users/telemetry", handlers.UsersGetTelemetry)
	mux.HandleFunc("GET /users/settings", handlers.UsersGetSettings)
	mux.HandleFunc("PUT /users/settings", handlers.UsersPutSettings)
	// Rotas de Auth
	mux.HandleFunc("POST /auth/login", handlers.Login)
	mux.HandleFunc("POST /auth/register", handlers.Register)
	mux.HandleFunc("POST /auth/logout", handlers.Logout)
}