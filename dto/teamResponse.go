package dto

type TeamResponse struct {
	Id   string `json:"team_id"`
	Name string `json:"name" db:"name"`
	AddressResponse
	CoachResponse
	PlayerResponse
}

type CoachResponse struct {
	Id           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	PhoneNumber  string `json:"phone_number"`
	EmailAddress string `json:"email_address"`
	AddressResponse
	Team string `json:"team_name"`
	Role string `json:"role"`
}
