package rest

import (
	"encoding/json"
	"net/http"

	"hackathon-pvc-backend/internal/registration/app"
	"hackathon-pvc-backend/internal/registration/domain"
)

type ParticipantHandler struct {
	service *app.ParticipantService
}

func NewParticipantHandler(service *app.ParticipantService) *ParticipantHandler {
	return &ParticipantHandler{
		service: service,
	}
}

func (h *ParticipantHandler) CreateParticipant(w http.ResponseWriter, r *http.Request) {
	var req CreateRegistrationRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "invalid JSON", err.Error())
		return
	}

	registration, err := h.service.Save(
		r.Context(),
		req.Name,
		req.Nickname,
		req.Email,
		req.Region,
		req.ProjectIdea,
		req.TeamPreference,
		req.DesiredTeammate,
	)
	if err != nil {
		h.writeErrorResponse(w, http.StatusBadRequest, "registration failed", err.Error())
		return
	}

	response := h.toRegistrationResponse(registration, req)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    response,
		"message": "Registration created successfully",
	})
}

func (h *ParticipantHandler) toRegistrationResponse(reg *domain.Participant, req CreateRegistrationRequest) CreateRegisterResponse {
	return CreateRegisterResponse{
		ID:              reg.ID(),
		Name:            reg.Name(),
		Nickname:        reg.Nickname(),
		Email:           reg.Email(),
		Region:          reg.Region(),
		ProjectIdea:     reg.ProjectIdea(),
		TeamPreference:  reg.TeamPreference(),
		DesiredTeammate: reg.DesiredTeammate(),
		RoleIDs:         req.RoleIDs,
		TechnologyIDs:   req.TechnologyIDs,
		CreatedAt:       reg.CreatedAt().Format("2006-01-02T15:04:05Z07:00"),
	}
}

func (h *ParticipantHandler) writeErrorResponse(w http.ResponseWriter, statusCode int, error, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error:   error,
		Message: message,
	})
}
