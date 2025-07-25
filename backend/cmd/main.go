package main

import (
	"log"
	"net/http"

	"hackathon-pvc-backend/cmd/config"
)

func main() {
	db, err := config.InitDatabase()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	deps := config.WireDependencies(db)

	router := config.SetupRouter(deps.RegistrationHandler)

	port := config.GetPort()
	log.Printf("🚀 Hackathon PVC API server starting on port %s", port)
	log.Printf("📊 Health check: http://localhost:%s/api/health", port)
	log.Printf("📝 Registration API: http://localhost:%s/api/registrations", port)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
