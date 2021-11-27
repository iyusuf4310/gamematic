package dto

import "k/golang/gamematic/errs"

type NewTeamRequest struct {
	Id   string `json:"team_id"`
	Name string `json:"name" db:"name"`
	AddressRequest
}

func (r NewTeamRequest) Validate() *errs.AppError {

	if len(r.Name) == 0 {
		return errs.NewValidationError("To create a new team you need to provide team name.")
	}
	return nil
}
