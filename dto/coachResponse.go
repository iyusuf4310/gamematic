package dto

type CoachResponse struct {
	Id           string `json:"id" db:"id"`
	FirstName    string `json:"coach.first_name"`
	LastName     string `json:"coach.last_name"`
	Gender       string `json:"coach.gender"`
	PhoneNumber  string `json:"coach.phone_number"`
	EmailAddress string `json:"coach.email_address"`
	AddressResponse
	Role string `json:"coach.role"`
	Team string `json:"team" db:"team"`
}
