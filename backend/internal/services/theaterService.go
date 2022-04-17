package services

import (
	"backend/internal/domain"
	"backend/internal/models"
)

type TheaterService struct {
	TheaterRepo domain.TheaterRepository
}

func NewTheaterService(repo domain.TheaterRepository) domain.TheaterService {
	return &TheaterService{
		TheaterRepo: repo,
	}
}

// CreateTheater implements domain.TheaterRepository
func (s *TheaterService) CreateTheater(theater *models.Theater) error {
	if err := s.TheaterRepo.CreateTheater(theater); err != nil {
		return err
	}

	return nil
}

// DeleteTheater implements domain.TheaterRepository
func (*TheaterService) DeleteTheater(id int) error {
	panic("unimplemented")
}

// GetAllTheater implements domain.TheaterRepository
func (s *TheaterService) GetAllTheater() ([]models.Theater, error) {
	lists, err := s.TheaterRepo.GetAllTheater()

	if err != nil {
		return nil, err
	}

	return lists, nil
}

// GetTheaterById implements domain.TheaterRepository
func (*TheaterService) GetTheaterById(id int) (models.Theater, error) {
	panic("unimplemented")
}

// UpdateTheater implements domain.TheaterRepository
func (*TheaterService) UpdateTheater(id int, theater *models.Theater) error {
	panic("unimplemented")
}
