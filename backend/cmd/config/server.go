package config

import (
	"net/http"
	"os"
	"time"

	"hackathon-pvc-backend/internal/registration/infrastructure/rest"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func SetupRouter(registrationHandler *rest.RegistrationHandler) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Route("/api", func(r chi.Router) {
		r.Get("/health", healthHandler)

		r.Route("/registrations", func(r chi.Router) {
			r.Post("/", registrationHandler.CreateRegistration)
		})

		r.Route("/html", func(r chi.Router) {
			r.Get("/", healthHtmlRenderHandler)
		})
	})

	return router
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok", "message": "Hackathon PVC API is running - ElWazy Pro"}`))
}

func healthHtmlRenderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Hackathon PVC API Health</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            margin: 0;
            padding: 0;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
        }
        .container {
            background: white;
            padding: 3rem;
            border-radius: 20px;
            box-shadow: 0 20px 40px rgba(0,0,0,0.1);
            text-align: center;
            max-width: 500px;
            width: 90%;
        }
        .status {
            font-size: 4rem;
            margin-bottom: 1rem;
        }
        .title {
            font-size: 2rem;
            color: #333;
            margin-bottom: 0.5rem;
            font-weight: 600;
        }
        .subtitle {
            color: #666;
            font-size: 1.1rem;
            margin-bottom: 2rem;
        }
        .badge {
            background: #10B981;
            color: white;
            padding: 0.5rem 1.5rem;
            border-radius: 50px;
            font-weight: 500;
            display: inline-block;
            margin-top: 1rem;
        }
        .timestamp {
            color: #999;
            font-size: 0.9rem;
            margin-top: 1rem;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="status">🚀</div>
        <h1 class="title">Hackathon PVC API</h1>
        <p class="subtitle">ElWazy Pro Edition</p>
        <div class="badge">✅ RUNNING</div>
        <p class="timestamp">Status checked at: ` + time.Now().Format("2006-01-02 15:04:05 MST") + `</p>
    </div>
</body>
</html>`

	w.Write([]byte(html))
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
