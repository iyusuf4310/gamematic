package service

import (
	"k/golang/gamematic/domain"
	"k/golang/gamematic/dto"
	"k/golang/gamematic/errs"
)

//go:generate mockgen -destination=../mocks/service/coachService.go -package=service k/golang/gamematic/service CoachService
type CoachService interface {
	GetAllCoaches() ([]dto.CoachResponse, *errs.AppError)
	GetCoach(name string) (*dto.CoachResponse, *errs.AppError)
	NewCoach(dto.NewCoachRequest) (*dto.CoachResponse, *errs.AppError)
}

type DefaultCoachService struct {
	repo domain.CoachRepository
}

func (s DefaultCoachService) GetAllCoaches() ([]dto.CoachResponse, *errs.AppError) {
	coach, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	response := make([]dto.CoachResponse, 0)
	for _, c := range coach {
		response = append(response, c.ToDTO())
	}
	return response, err
}

func (s DefaultCoachService) GetCoach(id string) (*dto.CoachResponse, *errs.AppError) {
	coach, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := coach.ToDTO()
	return &response, nil
}

func (s DefaultCoachService) NewCoach(r dto.NewCoachRequest) (*dto.CoachResponse, *errs.AppError) {
	err := r.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.Coach{
		Id:           r.Id,
		FirstName:    r.FirstName,
		LastName:     r.LastName,
		Gender:       r.Gender,
		PhoneNumber:  r.PhoneNumber,
		EmailAddress: r.EmailAddress,
		Address: domain.Address{
			Address1: r.AddressRequest.Address1,
			Address2: r.AddressRequest.Address2,
			City:     r.AddressRequest.City,
			State:    r.AddressRequest.State,
			Zipcode:  r.AddressRequest.Zipcode,
		},
		Role: r.Role,
		Team: r.Team,
	}
	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	response := newAccount.ToDTO()
	return &response, nil
}

func NewCoachService(repositery domain.CoachRepository) DefaultCoachService {
	return DefaultCoachService{repo: repositery}
}
