package rest

import (
	"encoding/json"
	"fmt"
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

func (h *RegistrationHandler) CreateRegistration(w http.ResponseWriter, r *http.Request) *hexttp.HTTPResponse {

	var req CreateRegistrationRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return hexttp.InvalidJSON()
	}

	fmt.Println("req")

	registration, err := h.service.Register(r.Context(), req.Name, req.Nickname, req.Email, req.Region, req.ProjectIdea, req.TeamPreference, req.DesiredTeammate, req.RoleIDs, req.TechnologyIDs)

	fmt.Println("registration")
	if err != nil {
		fmt.Println("Error creating registration:", err)
		return hexttp.InternalError("Error creating registration")
	}

	roles := make([]CreateRoleResponse, len(registration.Roles()))
	for i, role := range registration.Roles() {
		roles[i] = CreateRoleResponse{
			ID:          role.ID(),
			Name:        role.Name(),
			Description: role.Description(),
		}
	}

	technologies := make([]CreateTechnologyResponse, len(registration.Technologies()))
	for i, technology := range registration.Technologies() {
		technologies[i] = CreateTechnologyResponse{
			ID:          technology.ID(),
			Name:        technology.Name(),
			Description: technology.Description(),
		}
	}

	response := CreateRegistrationResponse{
		ID:              registration.ID(),
		Name:            registration.Participant().Name(),
		Email:           registration.Participant().Email(),
		Nickname:        registration.Participant().Nickname(),
		Region:          registration.Participant().Region(),
		ProjectIdea:     registration.Participant().ProjectIdea(),
		TeamPreference:  registration.Participant().TeamPreference(),
		DesiredTeammate: registration.Participant().DesiredTeammate(),
		RoleIDs:         roles,
		TechnologyIDs:   technologies,
		CreatedAt:       registration.Participant().CreatedAt().Format("2006-01-02T15:04:05Z07:00"),
	}

	return hexttp.Created(response)
}
