package domain

import (
	"k/golang/gamematic/errs"
	"k/golang/gamematic/logger"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"strings"
)

type CoachRepositoryDB struct {
	client *sqlx.DB
}

func (c CoachRepositoryDB) ById(id string) (*Coach, *errs.AppError) {
	var p Coach
	findCoachSql := "select * from coaches where coach_id = ?"
	err := c.client.Get(&p, findCoachSql, id)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, errs.NewNotFoundError("No Coach found with Id: " + id)
		} else {
			logger.Error("Error while querying coaches table " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected error")
		}
	}

	return &p, nil
}

func (c CoachRepositoryDB) Save(a Coach) (*Coach, *errs.AppError) {

	sqlInsert := "INSERT INTO coaches (first_name, last_name, gender, phone_number, email_address, address_1, address_2, city, state, zip_code, role, team) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"
	result, err := c.client.Exec(sqlInsert, a.FirstName, a.LastName, a.Gender, a.PhoneNumber, a.EmailAddress, a.Address.Address1, a.Address.Address2, a.Address.City, a.Address.State, a.Address.Zipcode, a.Role, a.Team)

	if err != nil {
		logger.Error("Error while creating new coach: " + err.Error())
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

func (p CoachRepositoryDB) FindAll() ([]Coach, *errs.AppError) {
	coaches := make([]Coach, 0)
	findAllSql := "select * from coaches"
	err := p.client.Select(&coaches, findAllSql)
	if err != nil {
		logger.Error("Error while querying coaches table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error")
	}
	return coaches, nil
}

func NewCoachRepositoryDB(dbClient *sqlx.DB) CoachRepositoryDB {
	return CoachRepositoryDB{dbClient}
}
