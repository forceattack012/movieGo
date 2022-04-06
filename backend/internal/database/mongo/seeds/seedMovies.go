package seeds

import "backend/internal/models"

func GetMovies() []models.Movie {
	listMovies := make([]models.Movie, 0)
	listMovies = append(listMovies, models.Movie{
		ID:        "4",
		Name:      "Spider-Man 1",
		Actors:    []string{"toby", "tom"},
		Directors: []string{"John Carpeter"},
		IMDB:      "imdb",
		Youtube:   "youtube",
		Duration:  120,
	},
	)

	listMovies = append(listMovies, models.Movie{
		ID:        "3",
		Name:      "Spider-Man 2",
		Actors:    []string{"toby", "tom", "andy"},
		Directors: []string{"John Carpeter", "johnathan jo start"},
		IMDB:      "imdb",
		Youtube:   "youtube",
		Duration:  300,
	},
	)

	return listMovies
}
