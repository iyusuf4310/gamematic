package service

import (
	"k/golang/gamematic/domain"
	"k/golang/gamematic/dto"

	"k/golang/gamematic/errs"
)

type PlayerService interface {
	GetAllPlayers() ([]domain.Player, *errs.AppError)
	GetPlayer(string) (*dto.PlayerResponse, *errs.AppError)
	NewPlayer(dto.NewPlayerRequest) (*dto.PlayerResponse, *errs.AppError)
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

func (s DefaultPlayerService) NewPlayer(r dto.NewPlayerRequest) (*dto.PlayerResponse, *errs.AppError) {

	err := r.Validate()
	if err != nil {
		return nil, err
	}

	a := domain.Player{
		Id:           r.Id,
		FirstName:    r.FirstName,
		LastName:     r.LastName,
		DateofBirth:  r.DateofBirth,
		Gender:       r.Gender,
		PhoneNumber:  r.PhoneNumber,
		EmailAddress: r.EmailAddress,
		JerseNumber:  r.JerseNumber,
		Team:         r.Team,
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

	response := s.ToDTO(newAccount)

	return &response, nil
}

func (s DefaultPlayerService) ToDTO(c *domain.Player) dto.PlayerResponse {
	return dto.PlayerResponse{
		Id:           c.Id,
		FirstName:    c.FirstName,
		LastName:     c.LastName,
		DateofBirth:  c.DateofBirth,
		Gender:       c.Gender,
		PhoneNumber:  c.PhoneNumber,
		EmailAddress: c.EmailAddress,
		JerseNumber:  c.JerseNumber,
		Team:         c.Team,
		AddressResponse: dto.AddressResponse{
			Address1: c.Address1,
			Address2: c.Address2,
			City:     c.City,
			State:    c.State,
			Zipcode:  c.Zipcode,
		},
	}
}

func (s DefaultPlayerService) GetAllPlayers() ([]domain.Player, *errs.AppError) {

	return s.repo.FindAll()
}

func NewPlayerService(repositery domain.PlayerRepository) DefaultPlayerService {
	return DefaultPlayerService{repo: repositery}
}
