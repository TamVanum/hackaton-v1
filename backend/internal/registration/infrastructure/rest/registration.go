package rest

import (
	"encoding/json"
	"net/http"

	"hackathon-pvc-backend/internal/registration/app"
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": "true",
		"data":    "ruta en proceso, tu tranqui que sale pronto esto",
		"message": "wena comparitooooo, aqui trabajando todavia y tu?",
	})
}

// func (h *ParticipantHandler) toRegistrationResponse(reg *domain.Participant, req CreateRegistrationRequest) CreateRegisterResponse {
// 	return CreateRegisterResponse{
// 		ID:              reg.ID(),
// 		Name:            reg.Name(),
// 		Nickname:        reg.Nickname(),
// 		Email:           reg.Email(),
// 		Region:          reg.Region(),
// 		ProjectIdea:     reg.ProjectIdea(),
// 		TeamPreference:  reg.TeamPreference(),
// 		DesiredTeammate: reg.DesiredTeammate(),
// 		RoleIDs:         req.RoleIDs,
// 		TechnologyIDs:   req.TechnologyIDs,
// 		CreatedAt:       reg.CreatedAt().Format("2006-01-02T15:04:05Z07:00"),
// 	}
// }

// func (h *ParticipantHandler) writeErrorResponse(w http.ResponseWriter, statusCode int, error, message string) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(statusCode)
// 	json.NewEncoder(w).Encode(ErrorResponse{
// 		Error:   error,
// 		Message: message,
// 	})
// }
