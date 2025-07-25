package rest

import (
	"encoding/json"
	"net/http"

	"hackathon-pvc-backend/internal/registration/application/services"
	"hackathon-pvc-backend/internal/registration/domain"
)

type RegistrationHandler struct {
	service *services.RegistrationService
}

func NewRegistrationHandler(service *services.RegistrationService) *RegistrationHandler {
	return &RegistrationHandler{
		service: service,
	}
}

func (h *RegistrationHandler) CreateRegistration(w http.ResponseWriter, r *http.Request) {
	var req CreateRegistrationRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "invalid JSON", err.Error())
		return
	}

	registration, err := h.service.RegisterParticipant(
		r.Context(),
		req.Name,
		req.Nickname,
		req.ProjectIdea,
		req.Teammate,
		req.Role,
	)
	if err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "registration failed", err.Error())
		return
	}

	// Convert to response DTO
	response := h.toRegistrationResponse(registration)

	// Write success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    response,
		"message": "Registration created successfully",
	})
}

func (h *RegistrationHandler) toRegistrationResponse(reg *domain.Registration) RegistrationResponse {
	return RegistrationResponse{
		ID:          reg.ID().Value(),
		Name:        reg.Name().Value(),
		Nickname:    reg.Nickname().Value(),
		ProjectIdea: reg.ProjectIdea().Value(),
		Teammate:    reg.Teammate().Value(),
		Role:        reg.Role().Value(),
		CreatedAt:   reg.CreatedAt().String(),
	}
}

func (h *RegistrationHandler) writeErrorResponse(w http.ResponseWriter, statusCode int, error, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error:   error,
		Message: message,
	})
}
