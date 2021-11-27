package service

import (
	"k/golang/gamematic/domain"
	"k/golang/gamematic/dto"
	"k/golang/gamematic/errs"
)

type TeamService interface {
	GetAllTeams() ([]dto.TeamResponse, *errs.AppError)
	GetTeam(name string) (*dto.TeamResponse, *errs.AppError)
	NewTeam(dto.NewTeamRequest) (*dto.TeamResponse, *errs.AppError)
}

type DefaultTeamService struct {
	repo domain.TeamRepository
}

func (s DefaultTeamService) GetAllTeams() ([]dto.TeamResponse, *errs.AppError) {

	team, err := s.repo.FindAll()

	if err != nil {
		return nil, err
	}

	response := make([]dto.TeamResponse, 0)

	for _, c := range team {
		response = append(response, c.ToDTO())
	}
	return response, err
}

func (s DefaultTeamService) GetTeam(name string) (*dto.TeamResponse, *errs.AppError) {

	team, err := s.repo.ByName(name)

	if err != nil {
		return nil, err
	}

	response := team.ToDTO()

	return &response, nil
}
func (s DefaultTeamService) NewTeam(r dto.NewTeamRequest) (*dto.TeamResponse, *errs.AppError) {

	err := r.Validate()
	if err != nil {
		return nil, err
	}

	a := domain.Team{
		Id:   r.Id,
		Name: r.Name,
		Address: domain.Address{
			Address1: r.AddressRequest.Address1,
			Address2: r.AddressRequest.Address2,
			City:     r.AddressRequest.City,
			State:    r.AddressRequest.State,
			Zipcode:  r.AddressRequest.Zipcode,
		},
	}

	newAccount, err := s.repo.Save(a)

	if err != nil {
		return nil, err
	}

	response := newAccount.ToDTO()

	return &response, nil
}

func NewTeamService(repositery domain.TeamRepository) DefaultTeamService {
	return DefaultTeamService{repo: repositery}
}
