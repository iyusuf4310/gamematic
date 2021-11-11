package domain

import "g/go/allsports/errs"

type Player struct {
	Id           string `json:"Id" db:"player_id"`
	FirstName    string `json:"FirstName" db:"first_name"`
	LastName     string `json:"LastName" db:"last_name"`
	DateofBirth  string `json:"DateofBirth" db:"birth_date"`
	Gender       string `json:"Gender" db:"gender"`
	PhoneNumber  string `json:"PhoneNumber" db:"phone_number"`
	EmailAddress string `json:"EmailAddress" db:"email_address"`
	JerseNumber  string `json:"JerseNumber" db:"jerse_number"`
	Team         string `json:"Team" db:"team"`
	Address
}

type PlayerRepository interface {
	FindAll() ([]Player, *errs.AppError)
	ById(string) (*Player, *errs.AppError)
}
