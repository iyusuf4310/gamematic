package dto

import "k/golang/gamematic/errs"

type NewPlayerRequest struct {
	Id           string `json:"player_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	DateofBirth  string `json:"birth_date"`
	Gender       string `json:"gender"`
	PhoneNumber  string `json:"phone_number"`
	EmailAddress string `json:"email_address"`
	JerseNumber  string `json:"jerse_number"`
	Team         string `json:"team"`
	AddressRequest
}

type AddressRequest struct {
	Address1 string `json:"address_1"`
	Address2 string `json:"address_2"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zipcode  string `json:"zip_code"`
}

func (r NewPlayerRequest) Validate() *errs.AppError {

	if len(r.Team) == 0 {
		return errs.NewValidationError("To create a new player you need to provide team name.")
	}
	return nil
}
