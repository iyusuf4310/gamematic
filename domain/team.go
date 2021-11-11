package domain

import "g/go/allsports/errs"

type Team struct {
	Id   string `json:"team_id" db:"team_id"`
	Name string `json:"name" db:"name"`
	Address
	Coach
	Player
}

type Address struct {
	Address1 string `json:"Address1" db:"address_1"`
	Address2 string `json:"Address2" db:"address_2"`
	City     string `json:"City" db:"city"`
	State    string `json:"State" db:"state"`
	Zipcode  string `json:"Zipcode" db:"zip_code"`
}

type Coach struct {
	Id           string `json:"id" db:"id"`
	FirstName    string `json:"first_name" db:"first_name"`
	LastName     string `json:"last_name" db:"last_name"`
	Gender       string `json:"gender" db:"gender"`
	PhoneNumber  string `json:"phone_number" db:"phone_number"`
	EmailAddress string `json:"email_address" db:"email_address"`
	Address      Address
	Team         string `json:"team_name" db:"team"`
	Role         string `json:"role" db:"role"`
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
}
