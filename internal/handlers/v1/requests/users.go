package requests

type UserSignUpRequest struct {
	Username string `json:"username" validate:"required" example:"john_doe"`
}

type UserUpdateRequest struct {
	Id            uint     `json:"id"`
	FirstName     string   `json:"first_name" example:"John"`
	LastName      string   `json:"last_name" example:"Doe"`
	Position      string   `json:"position" example:"Frontend Developer"`
	Qualification string   `json:"qualification" example:"Middle"`
	About         string   `json:"about" example:"Brief description"`
	Skillset      []string `json:"skillset" example:"Javascript,Typescript,Vue"`
}

type UserDeleteRequest struct {
	Id uint `json:"id" validate:"required" example:"1"`
}
