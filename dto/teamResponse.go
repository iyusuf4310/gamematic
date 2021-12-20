package dto

type TeamResponse struct {
	Id   string `json:"team_id"`
	Name string `json:"name" db:"name"`
	AddressResponse
}
