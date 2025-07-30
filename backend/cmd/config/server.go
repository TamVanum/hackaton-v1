package config

import (
	"net/http"
	"os"

	"hackathon-pvc-backend/internal/registration/infrastructure/rest"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/tamvanum/go-hexttp/hexttp"
)

func SetupRouter(registrationHandler *rest.RegistrationHandler, rolesHandler *rest.RolesHandler) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Route("/api", func(r chi.Router) {
		r.Get("/health", healthHandler)

		r.Route("/registrations", func(r chi.Router) {
			r.Post("/", hexttp.Make(registrationHandler.Create))
		})

		r.Route("/roles", func(r chi.Router) {
			r.Get("/", hexttp.Make(rolesHandler.RetrieveRoles))
			r.Post("/", hexttp.Make(rolesHandler.CreateRole))
		})
	})

	return router
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok", "message": "Hackathon PVC API is running"}`))
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
