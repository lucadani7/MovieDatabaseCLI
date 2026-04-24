package models

// MovieResponse matches the top-level JSON object returned by TMDB
type MovieResponse struct {
	Page    int     `json:"page"`
	Results []Movie `json:"results"`
}

// Movie matches the individual movie objects inside the "results" array
type Movie struct {
	Title       string  `json:"title"`
	ReleaseDate string  `json:"release_date"`
	VoteAverage float64 `json:"vote_average"`
	Overview    string  `json:"overview"`
}
