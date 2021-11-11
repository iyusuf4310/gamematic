package service

import (
	"k/golang/gamematic/domain"
	"k/golang/gamematic/dto"
	"k/golang/gamematic/errs"
)

type PlayerService interface {
	GetAllPlayers() ([]domain.Player, *errs.AppError)
	GetPlayer(string) (*dto.PlayerResponse, *errs.AppError)
}

type DefaultPlayerService struct {
	repo domain.PlayerRepository
}

func (s DefaultPlayerService) GetPlayer(id string) (*dto.PlayerResponse, *errs.AppError) {

	p, err := s.repo.ById(id)

	if err != nil {
		return nil, err
	}

	response := s.ToDTO(p)

	return &response, nil
}

func (s DefaultPlayerService) ToDTO(c *domain.Player) dto.PlayerResponse {
	return dto.PlayerResponse{
		Id:              c.Id,
		FirstName:       c.FirstName,
		LastName:        c.LastName,
		DateofBirth:     c.DateofBirth,
		Gender:          c.Gender,
		PhoneNumber:     c.PhoneNumber,
		EmailAddress:    c.EmailAddress,
		JerseNumber:     c.JerseNumber,
		Team:            c.Team,
		AddressResponse: dto.AddressResponse(c.Address),
	}
}

func (s DefaultPlayerService) GetAllPlayers() ([]domain.Player, *errs.AppError) {

	return s.repo.FindAll()
}

func NewPlayerService(repositery domain.PlayerRepository) DefaultPlayerService {
	return DefaultPlayerService{repo: repositery}
}
