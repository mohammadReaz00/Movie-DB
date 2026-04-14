package handlers

import (
	"encoding/json"
	"net/http"

	"mohammadreaz.com/moviedb/models"
)

type MovieHandler struct{}

func (h *MovieHandler) writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{ID: 1,
			TMDB_ID:     100,
			Title:       "The Hacker",
			ReleaseYear: 2023,
			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 1, FirstName: "Max"}},
		},
		{ID: 2,
			TMDB_ID:     101,
			Title:       "The Hacker 2",
			ReleaseYear: 2025,
			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 1, FirstName: "Max"}},
		},
	}

	h.writeJSONResponse(w, movies)

}

func (h *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{ID: 1,
			TMDB_ID:     100,
			Title:       "The Hacker from Rnadom",
			ReleaseYear: 2023,
			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 1, FirstName: "Max"}},
		},
		{ID: 2,
			TMDB_ID:     101,
			Title:       "The Hacker from Random 2",
			ReleaseYear: 2025,
			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 1, FirstName: "Max"}},
		},
	}

	h.writeJSONResponse(w, movies)

}
