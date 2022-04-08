package repositories

import (
	"backend/internal/domain"
	"backend/internal/models"
	"context"
	"log"

	"github.com/naamancurtis/mongo-go-struct-to-bson/mapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type movieRepository struct {
	MongoDB    *mongo.Database
	collection string
}

func NewMovieRepository(database *mongo.Database, c string) domain.MovieRepository {
	return &movieRepository{
		MongoDB:    database,
		collection: c,
	}
}

func (m *movieRepository) GetAllMovies(ctx context.Context) ([]models.Movie, error) {
	query, error := m.MongoDB.Collection(m.collection).Find(ctx, bson.D{})

	if error != nil {
		log.Println("error", error)
		return nil, error
	}

	defer query.Close(ctx)

	listMovies := make([]models.Movie, 0)

	for query.Next(ctx) {
		var row models.Movie
		err := query.Decode(&row)

		if err != nil {
			log.Println("error", err)
		}

		listMovies = append(listMovies, row)
	}

	return listMovies, nil
}

// CreateMovie implements domain.MovieRepository
func (m *movieRepository) CreateMovie(ctx context.Context, Movie *models.Movie) error {
	result := mapper.ConvertStructToBSONMap(Movie, nil)
	res, err := m.MongoDB.Collection(m.collection).InsertOne(ctx, result)

	if err != nil {
		log.Fatal(err, res)
	}

	return nil
}

// DeleteMovie implements domain.MovieRepository
func (m *movieRepository) DeleteMovie(ctx context.Context, id string) (int64, error) {
	filter := bson.D{primitive.E{Key: "ID", Value: id}}
	result, err := m.MongoDB.Collection(m.collection).DeleteOne(ctx, filter)

	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return result.DeletedCount, nil
}

// GetMovieById implements domain.MovieRepository
func (m *movieRepository) GetMovieById(ctx context.Context, id string) (models.Movie, error) {
	result := models.Movie{}
	filter := bson.D{primitive.E{Key: "ID", Value: id}}

	err := m.MongoDB.Collection(m.collection).FindOne(ctx, filter).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}

// UpdateMovie implements domain.MovieRepository
func (m *movieRepository) UpdateMovie(ctx context.Context, id string, movie *models.Movie) (int64, error) {
	filter := bson.D{primitive.E{Key: "ID", Value: id}}

	update := bson.M{"$set": bson.M{
		"Name":      movie.Name,
		"Actors":    movie.Actors,
		"Directors": movie.Directors,
		"IMDB":      movie.IMDB,
		"Youtube":   movie.Youtube,
		"Duration":  movie.Duration,
		"StartDate": movie.StartDate,
		"Synopsis":  movie.Synopsis,
		"Image":     movie.Image}}

	result, err := m.MongoDB.Collection(m.collection).UpdateOne(ctx, filter, update)

	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return result.ModifiedCount, err
}
