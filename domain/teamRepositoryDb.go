package domain

import (
	"database/sql"
	"g/go/allsports/errs"
	"g/go/allsports/logger"
	"time"

	"github.com/jmoiron/sqlx"
)

type TeamRepositoryDb struct {
	client *sqlx.DB
}

func (team TeamRepositoryDb) ByName(name string) (*Team, *errs.AppError) {
	var t Team

	findTeamSql := "select * from teams t join coaches c join players p where t.name = ? and t.name = p.team and c.team = t.name"

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

	findAllSql := "select * from teams t join coaches c join players p where t.name = p.team and c.team = t.name"

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

func NewTeamRepositoryDb() TeamRepositoryDb {
	client, err := sqlx.Open("mysql", "root:dayaxQ!@tcp(localhost:3306)/soccer")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return TeamRepositoryDb{client}
}
