package dto

import "k/golang/gamematic/errs"

type NewCoachRequest struct {
	Id           string `json:"id" db:"id"`
	FirstName    string `json:"first_name" db:"first_name"`
	LastName     string `json:"last_name" db:"last_name"`
	Gender       string `json:"gender" db:"gender"`
	PhoneNumber  string `json:"phone_number" db:"phone_number"`
	EmailAddress string `json:"email_address" db:"email_address"`
	AddressRequest
	Role string `json:"role" db:"role"`
	Team string `json:"team" db:"team"`
}

func (r NewCoachRequest) Validate() *errs.AppError {

	if len(r.Team) == 0 {
		return errs.NewValidationError("To create a new coach you need to provide team name.")
	}
	return nil
}
