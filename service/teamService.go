package service

import (
	"g/go/allsports/domain"
	"g/go/allsports/errs"
)

type TeamService interface {
	GetAllTeams() ([]domain.Team, *errs.AppError)
	GetTeam(name string) (*domain.Team, *errs.AppError)
}

type DefaultTeamService struct {
	repo domain.TeamRepository
}

func (s DefaultTeamService) GetAllTeams() ([]domain.Team, *errs.AppError) {

	return s.repo.FindAll()
}

func (s DefaultTeamService) GetTeam(name string) (*domain.Team, *errs.AppError) {

	return s.repo.ByName(name)
}

func NewTeamService(repositery domain.TeamRepository) DefaultTeamService {
	return DefaultTeamService{repo: repositery}
}
