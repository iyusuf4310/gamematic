package domain

import (
	"k/golang/gamematic/dto"
	"k/golang/gamematic/errs"
)

type Coach struct {
	Id           string `json:"id" db:"coach_id"`
	FirstName    string `json:"first_name" db:"first_name"`
	LastName     string `json:"last_name" db:"last_name"`
	Gender       string `json:"gender" db:"gender"`
	PhoneNumber  string `json:"phone_number" db:"phone_number"`
	EmailAddress string `json:"email_address" db:"email_address"`
	Address
	Role string `json:"role" db:"role"`
	Team string `json:"team_name" db:"team"`
}

type CoachRepository interface {
	FindAll() ([]Coach, *errs.AppError)
	ById(name string) (*Coach, *errs.AppError)
	Save(c Coach) (*Coach, *errs.AppError)
}

func (c Coach) ToDTO() dto.CoachResponse {
	return dto.CoachResponse{
		Id:           c.Id,
		FirstName:    c.FirstName,
		LastName:     c.LastName,
		Gender:       c.Gender,
		PhoneNumber:  c.PhoneNumber,
		EmailAddress: c.EmailAddress,
		AddressResponse: dto.AddressResponse{
			Address1: c.Address.Address1,
			Address2: c.Address.Address2,
			City:     c.Address.City,
			State:    c.Address.State,
			Zipcode:  c.Address.Zipcode,
		},
		Role: c.Role,
		Team: c.Team,
	}
}
