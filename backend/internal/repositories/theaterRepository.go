package repositories

import (
	"backend/internal/domain"
	"backend/internal/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type TheaterRepository struct {
	MongoDB    *mongo.Database
	Collection string
}

func NewTheaterRepository(db *mongo.Database, collection string) domain.TheaterRepository {
	return &TheaterRepository{
		MongoDB:    db,
		Collection: collection,
	}
}

// CreateTheater implements domain.TheaterRepository
func (*TheaterRepository) CreateTheater(ctx context.Context, theater *models.Theater) error {
	panic("unimplemented")
}

// DeleteTheater implements domain.TheaterRepository
func (*TheaterRepository) DeleteTheater(ctx context.Context, id int) error {
	panic("unimplemented")
}

// GetAllTheater implements domain.TheaterRepository
func (*TheaterRepository) GetAllTheater(ctx context.Context) ([]models.Theater, error) {
	panic("unimplemented")
}

// GetTheaterById implements domain.TheaterRepository
func (*TheaterRepository) GetTheaterById(ctx context.Context, id int) (models.Theater, error) {
	panic("unimplemented")
}

// UpdateTheater implements domain.TheaterRepository
func (*TheaterRepository) UpdateTheater(ctx context.Context, id int, theater *models.Theater) error {
	panic("unimplemented")
}
