package domain

import (
	"k/golang/gamematic/errs"
	"k/golang/gamematic/logger"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type PlayerRepositoryDb struct {
	client *sqlx.DB
}

func (pl PlayerRepositoryDb) ById(id string) (*Player, *errs.AppError) {
	var p Player
	findPlayerSql := "select * from players where player_id = ?"

	err := pl.client.Get(&p, findPlayerSql, id)

	if err != nil {
		logger.Error("Error while querying players table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error")
	}

	return &p, nil
}

func (p PlayerRepositoryDb) Save(a Player) (*Player, *errs.AppError) {
	sqlInsert := "INSERT INTO players (player_id, first_name, last_name, birth_date, gender, phone_number, email_address, jerse_number, team, address_1, address_2, city, state, zip_code) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	result, err := p.client.Exec(sqlInsert, a.Id, a.FirstName, a.FirstName, a.DateofBirth, a.Gender, a.PhoneNumber, a.EmailAddress, a.JerseNumber, a.Team, a.Address1, a.Address2, a.City, a.State, a.Zipcode)

	if err != nil {
		logger.Error("Error while creating new user: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert Id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error")
	}

	a.Id = strconv.FormatInt(id, 10)

	return &a, nil
}

func (p PlayerRepositoryDb) FindAll() ([]Player, *errs.AppError) {

	players := make([]Player, 0)

	findAllSql := "select * from players"

	err := p.client.Select(&players, findAllSql)

	if err != nil {
		logger.Error("Error while querying players table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error")
	}
	return players, nil
}

func NewPlayerRepositoryDb(dbClient *sqlx.DB) PlayerRepositoryDb {
	return PlayerRepositoryDb{dbClient}
}
