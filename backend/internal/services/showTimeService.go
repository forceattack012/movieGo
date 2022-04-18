package services

import (
	"backend/internal/domain"
	"backend/internal/models"
	"context"
	"fmt"
	"sort"
	"time"
)

type ShowTimeService struct {
	repo      domain.ShowTimeRepository
	movieRepo domain.MovieRepository
}

func NewShowTimeService(repo domain.ShowTimeRepository, mRepo domain.MovieRepository) domain.ShowTimeService {
	return &ShowTimeService{
		repo:      repo,
		movieRepo: mRepo,
	}
}

// GetShowTimesByMovieId implements domain.ShowTimeService
func (service *ShowTimeService) GetShowTimesByMovieId(movieId string) ([]models.ResponseShowTime, error) {
	lists, err := service.repo.GetShowTimesByMovieId(movieId)

	if err != nil {
		return nil, err
	}

	showTimes := make([]models.ResponseShowTime, 0)
	times := make(map[int64][]models.Times, 0)
	keys := make(map[int]bool)

	for _, l := range lists {
		times[int64(l.TheaterId)] = append(times[int64(l.Id)], models.Times{
			Id:   l.Id,
			Time: l.Time,
		})
	}

	for _, sh := range lists {
		if _, value := keys[int(sh.TheaterId)]; !value {
			keys[int(sh.TheaterId)] = true
			sort.Sort(models.TimeSlice(times[int64(sh.TheaterId)]))

			showTimes = append(showTimes, models.ResponseShowTime{
				Times:     times[int64(sh.TheaterId)],
				MovieId:   sh.MovieId,
				TheaterId: sh.TheaterId,
				Theater:   sh.Theater,
			})
		}
	}

	return showTimes, nil
}

// CreateShowTime implements domain.ShowTimeService
func (service *ShowTimeService) CreateShowTime(showtime *models.ShowTime) error {
	isShowtime, err := service.repo.IsShowTimeByTheaterAndTime(int64(showtime.TheaterId), showtime.Time)

	if isShowtime {
		return err
	}

	if err := service.repo.CreateShowTime(showtime); err != nil {
		return err
	}

	return nil
}

// DeleteShowTime implements domain.ShowTimeService
func (service *ShowTimeService) DeleteShowTime(id int64) error {
	showTime, err := service.repo.GetAllShowTimeById(id)

	if err != nil {
		return err
	}

	if err := service.repo.DeleteShowTime(id, showTime); err != nil {
		return err
	}

	return nil
}

// GetAllShowTime implements domain.ShowTimeService
func (service *ShowTimeService) GetAllShowTime() ([]models.ShowTime, error) {
	list, err := service.repo.GetAllShowTime()

	if err != nil {
		return nil, err
	}

	movies := make([]models.Movie, 0)
	for _, l := range list {
		movie, err := service.movieRepo.GetMovieById(context.Background(), l.MovieId)

		if err != nil {
			return nil, err
		}

		movies = append([]models.Movie{}, movie)
	}

	fmt.Println(movies[0].Name)
	return list, nil
}

// GetShowTimeSlot implements domain.ShowTimeService
func (*ShowTimeService) GetShowTimeSlot(theaterId int64) []time.Time {
	panic("unimplemented")
}
