package rest

import (
	"encoding/json"
	"net/http"

	"hackathon-pvc-backend/internal/registration/app"

	"github.com/tamvanum/go-hexttp/hexttp"
)

type RegistrationHandler struct {
	service *app.RegistrationService
}

func NewRegistrationHandler(service *app.RegistrationService) *RegistrationHandler {
	return &RegistrationHandler{
		service: service,
	}
}

func (h *RegistrationHandler) Create(w http.ResponseWriter, r *http.Request) *hexttp.HTTPResponse {

	var req CreateRegistrationRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return hexttp.InvalidJSON()
	}

	registration, err := h.service.Create(r.Context(), req.Name, req.Email, req.Nickname, req.Region, req.ProjectIdea, req.TeamPreference, req.DesiredTeammate, req.RoleIDs, req.TechnologyIDs)
	if err != nil {
		return hexttp.InternalError(err.Error())
	}

	return hexttp.Created(registration)
}
