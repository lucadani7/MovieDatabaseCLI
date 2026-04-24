package api

import (
	"MovieDatabaseCLI/models"
	"encoding/json"
	"fmt"
	"net/http"
)

// FetchMovies reaches out to TMDB and returns a parsed MovieResponse struct
func FetchMovies(endpoint string, apiKey string) (*models.MovieResponse, error) {
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%s?language=en-US&page=1", endpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}
	defer func() {
		if closeErr := res.Body.Close(); closeErr != nil {
			fmt.Printf("Warning: failed to close response body: %v\n", closeErr)
		}
	}()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status code: %d", res.StatusCode)
	}
	var movieResponse models.MovieResponse
	if err := json.NewDecoder(res.Body).Decode(&movieResponse); err != nil {
		return nil, fmt.Errorf("decoding JSON: %w", err)
	}
	return &movieResponse, nil
}
