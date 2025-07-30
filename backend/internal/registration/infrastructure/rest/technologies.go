package rest

import (
	"encoding/json"
	"hackathon-pvc-backend/internal/registration/app"
	"net/http"

	"github.com/tamvanum/go-hexttp/hexttp"
)

type TechnologiesHandler struct {
	service *app.TechnologyService
}

func NewTechnologiesHandler(service *app.TechnologyService) *TechnologiesHandler {
	return &TechnologiesHandler{
		service: service,
	}
}

func (h *TechnologiesHandler) RetrieveTechnologies(w http.ResponseWriter, r *http.Request) *hexttp.HTTPResponse {
	technologies, err := h.service.Get(r.Context())
	if err != nil {
		return hexttp.InternalError("Error getting technologies")
	}

	response := make([]CreateTechnologyResponse, len(technologies))
	for i, technology := range technologies {
		response[i] = CreateTechnologyResponse{
			ID:          technology.ID(),
			Name:        technology.Name(),
			Description: technology.Description(),
		}
	}

	return hexttp.OK(response)
}

func (h *TechnologiesHandler) CreateTechnology(w http.ResponseWriter, r *http.Request) *hexttp.HTTPResponse {
	var req CreateTechnologyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return hexttp.InvalidJSON()
	}

	technology, err := h.service.Create(r.Context(), req.Name, req.Description)
	if err != nil {
		return hexttp.InternalError("Error creating technology")
	}

	response := CreateTechnologyResponse{
		ID:          technology.ID(),
		Name:        technology.Name(),
		Description: technology.Description(),
	}

	return hexttp.Created(response)
}
