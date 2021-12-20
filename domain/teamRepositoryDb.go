package domain

import (
	"database/sql"
	"k/golang/gamematic/errs"
	"k/golang/gamematic/logger"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type TeamRepositoryDb struct {
	client *sqlx.DB
}

func (team TeamRepositoryDb) ByName(name string) (*Team, *errs.AppError) {
	var t Team

	findTeamSql := "select * from teams  where name = ?"

	err := team.client.Get(&t, findTeamSql, name)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Team Not Found: " + name)
		} else {
			logger.Error("Error while querying team table " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected Error")
		}
	}

	return &t, nil
}

func (tm TeamRepositoryDb) FindAll() ([]Team, *errs.AppError) {
	teams := make([]Team, 0)

	findAllSql := "select t.team_id, t.name, t.address_1, t.address_2, t.city, t.state, t.zip_code, c.* from teams t join coaches c where c.team = t.name"

	err := tm.client.Select(&teams, findAllSql)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Team Not Found: ")
		} else {
			logger.Error("Error while querying team table " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected Error")
		}
	}

	return teams, nil
}

func (t TeamRepositoryDb) Delete(ID int) *errs.AppError {

	deleteTeamSql := "DELETE from teams where team_id=?"

	res, err := t.client.Exec(deleteTeamSql, ID)

	count, _ := res.RowsAffected()

	if err != nil {
		if count > 0 {
			return errs.NewNotFoundError("Team Not Found: " + strconv.Itoa(ID))
		} else {
			logger.Error("Error while deleting team from team table " + err.Error())
			return errs.NewUnexpectedError("Unexpected Error")
		}
	}

	return nil
}

func (t TeamRepositoryDb) Save(a Team) (*Team, *errs.AppError) {
	sqlInsert := "INSERT INTO teams (name, address_1, address_2, city, state, zip_code) VALUES (?,?,?,?,?,?)"
	result, err := t.client.Exec(sqlInsert, a.Name, a.Address1, a.Address2, a.City, a.State, a.Zipcode)

	if err != nil {
		logger.Error("Error while creating new team: " + err.Error())
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

func (t TeamRepositoryDb) Update(a Team) (*Team, *errs.AppError) {
	a.Id = "1018"
	sqlInsert := "UPDATE teams SET name = ?, address_1 = ?, address_2 = ?, city = ?, state = ?, zip_code = ? where team_id = ? "
	_, err := t.client.Exec(sqlInsert, a.Name, a.Address1, a.Address2, a.City, a.State, a.Zipcode, a.Id)

	if err != nil {
		logger.Error("Error while Updating team: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error")
	}

	// id, err := result.LastInsertId()
	// if err != nil {
	// 	logger.Error("Error while getting last insert Id: " + err.Error())
	// 	return nil, errs.NewUnexpectedError("Unexpected error")
	// }



	return &a, nil
}

func NewTeamRepositoryDb(dbClient *sqlx.DB) TeamRepositoryDb {
	return TeamRepositoryDb{dbClient}
}
