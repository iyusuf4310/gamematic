package service

import (
	"k/golang/gamematic/domain"
	"k/golang/gamematic/errs"
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
