package domain

import (
	"k/golang/gamematic/dto"
	"k/golang/gamematic/errs"
)

type Team struct {
	Id   string `json:"team_id" db:"team_id"`
	Name string `json:"name" db:"name"`
	Address
	Coach
}

type Address struct {
	Address1 string `json:"Address1" db:"address_1"`
	Address2 string `json:"Address2" db:"address_2"`
	City     string `json:"City" db:"city"`
	State    string `json:"State" db:"state"`
	Zipcode  string `json:"Zipcode" db:"zip_code"`
}

type Assistant struct {
	Id          string  `json:"id"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Gender      string  `json:"gender"`
	PhoneNumber string  `json:"phone_number"`
	Address     Address `json:"address"`
}

type TeamRepository interface {
	FindAll() ([]Team, *errs.AppError)
	ByName(name string) (*Team, *errs.AppError)
	Save(t Team) (*Team, *errs.AppError)
}

func (c Team) ToDTO() dto.TeamResponse {
	return dto.TeamResponse{
		Id:   c.Id,
		Name: c.Name,
		AddressResponse: dto.AddressResponse{
			Address1: c.Address1,
			Address2: c.Address2,
			City:     c.City,
			State:    c.State,
			Zipcode:  c.Zipcode,
		},
		CoachResponse: dto.CoachResponse{
			FirstName:    c.FirstName,
			LastName:     c.LastName,
			Gender:       c.Gender,
			PhoneNumber:  c.PhoneNumber,
			EmailAddress: c.EmailAddress,
			Role:         c.Role,
			Team:         c.Team,
		},
	}
}
