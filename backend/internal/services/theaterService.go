package services

import (
	"backend/internal/domain"
	"backend/internal/models"
	"errors"
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
	if s.getSize(theater.Size) == -1 {
		return errors.New("Size Not Found")
	}

	theater.Seat = make([]models.Seat, 0)

	for i := 1; i <= s.getSize(theater.Size); i++ {
		theater.Seat = append(theater.Seat, models.Seat{
			SeatNumber: i,
			TheaterId:  int(theater.Id),
		})
	}

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

func (t *TheaterService) getSize(s string) int {
	m := map[string]int{"Small": 20, "Medium": 30, "Large": 60}
	val, result := m[s]

	if !result {
		return -1
	}
	return val
}
