package rest

type CreateRegistrationRequest struct {
	Name            string  `json:"name" binding:"required"`
	Email           string  `json:"email" binding:"required"`
	Nickname        string  `json:"nickname" binding:"required"`
	Region          string  `json:"region" binding:"required"`
	ProjectIdea     string  `json:"project_idea" binding:"required"`
	TeamPreference  bool    `json:"team_preference" binding:"required"`
	DesiredTeammate *string `json:"desired_teammate,omitempty"`
	Role            string  `json:"role" binding:"required"`
}

type RegistrationResponse struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Nickname        string  `json:"nickname"`
	Email           string  `json:"email"`
	Region          string  `json:"region"`
	ProjectIdea     string  `json:"project_idea"`
	TeamPreference  bool    `json:"team_preference"`
	DesiredTeammate *string `json:"desired_teammate,omitempty"`
	Role            string  `json:"role"`
	CreatedAt       string  `json:"created_at"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
