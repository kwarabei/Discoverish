package responses

type UserGetResponse struct {
	Id            uint     `json:"id"`
	Username      string   `json:"username"`
	FirstName     string   `json:"first_name"`
	LastName      string   `json:"last_name"`
	Position      string   `json:"position"`
	Qualification string   `json:"qualification"`
	About         string   `json:"about"`
	Skillset      []string `json:"skillset"`
}

type UserGetListResponse struct {
	Users []UserGetResponse `json:"users"`
}
