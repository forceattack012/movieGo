package repositories

import (
	"backend/internal/domain"
	"backend/internal/models"

	"gorm.io/gorm"
)

type TheaterRepository struct {
	db *gorm.DB
}

func NewTheaterRepository(db *gorm.DB) domain.TheaterRepository {
	return &TheaterRepository{
		db: db,
	}
}

// CreateTheater implements domain.TheaterRepository
func (t *TheaterRepository) CreateTheater(theater *models.Theater) error {
	if err := t.db.Create(theater).Error; err != nil {
		return err
	}

	return nil
}

// DeleteTheater implements domain.TheaterRepository
func (*TheaterRepository) DeleteTheater(id int) error {
	panic("unimplemented")
}

// GetAllTheater implements domain.TheaterRepository
func (t *TheaterRepository) GetAllTheater() ([]models.Theater, error) {
	theaters := make([]models.Theater, 0)

	if err := t.db.Preload("Seat").Find(&theaters).Error; err != nil {
		return nil, err
	}

	return theaters, nil
}

// GetTheaterById implements domain.TheaterRepository
func (*TheaterRepository) GetTheaterById(id int) (models.Theater, error) {
	panic("unimplemented")
}

// UpdateTheater implements domain.TheaterRepository
func (*TheaterRepository) UpdateTheater(id int, theater *models.Theater) error {
	panic("unimplemented")
}
