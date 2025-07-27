package rest

type CreateRegistrationRequest struct {
	Name            string  `json:"name" binding:"required"`
	Nickname        string  `json:"nickname" binding:"required"`
	ProjectIdea     string  `json:"project_idea" binding:"required"`
	DesiredTeammate *string `json:"desired_teammate,omitempty"`
	Role            string  `json:"role" binding:"required"`
}

type RegistrationResponse struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Nickname        string  `json:"nickname"`
	ProjectIdea     string  `json:"project_idea"`
	DesiredTeammate *string `json:"desired_teammate,omitempty"`
	Role            string  `json:"role"`
	CreatedAt       string  `json:"created_at"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
