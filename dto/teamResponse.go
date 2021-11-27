package dto

type TeamResponse struct {
	Id   string `json:"team_id"`
	Name string `json:"name" db:"name"`
	AddressResponse
	CoachResponse
}

type CoachResponse struct {
	FirstName    string `json:"-"`
	LastName     string `json:"-"`
	Gender       string `json:"-"`
	PhoneNumber  string `json:"-"`
	EmailAddress string `json:"-"`
	Role         string `json:"-"`
}
