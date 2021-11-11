package domain

import (
	"k/golang/gamematic/errs"
	"k/golang/gamematic/logger"
	"time"

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

func NewPlayerRepositoryDb() PlayerRepositoryDb {
	client, err := sqlx.Open("mysql", "root:dayaxQ!@tcp(localhost:3306)/soccer")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return PlayerRepositoryDb{client}
}
