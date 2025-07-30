package rest

import (
	"encoding/json"
	"hackathon-pvc-backend/internal/registration/app"
	"net/http"

	"github.com/tamvanum/go-hexttp/hexttp"
)

type RolesHandler struct {
	service *app.RoleService
}

func NewRolesHandler(service *app.RoleService) *RolesHandler {
	return &RolesHandler{
		service: service,
	}
}

func (h *RolesHandler) RetrieveRoles(w http.ResponseWriter, r *http.Request) *hexttp.HTTPResponse {
	roles, err := h.service.Get(r.Context())
	if err != nil {
		return hexttp.InternalError("Error getting roles")
	}

	response := make([]CreateRoleResponse, len(roles))
	for i, role := range roles {
		response[i] = CreateRoleResponse{
			ID:          role.ID(),
			Name:        role.Name(),
			Description: role.Description(),
		}
	}

	return hexttp.OK(response)

}

func (h *RolesHandler) CreateRole(w http.ResponseWriter, r *http.Request) *hexttp.HTTPResponse {
	var req CreateRoleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return hexttp.InvalidJSON()
	}

	role, err := h.service.Create(r.Context(), req.Name, req.Description)
	if err != nil {
		return hexttp.InternalError("Error creating role")
	}

	response := CreateRoleResponse{
		ID:          role.ID(),
		Name:        role.Name(),
		Description: role.Description(),
	}

	return hexttp.Created(response)
}
