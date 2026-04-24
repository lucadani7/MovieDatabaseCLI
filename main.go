package main

import (
	"MovieDatabaseCLI/api"
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/joho/godotenv"
	"github.com/pterm/pterm"
)

var CLI struct {
	Type string `help:"Type of movies to fetch." enum:"playing,popular,top,upcoming" default:"popular" short:"t"`
}

func main() {
	_ = godotenv.Load()
	kong.Parse(&CLI)
	endpoints := map[string]string{
		"playing":  "now_playing",
		"popular":  "popular",
		"top":      "top_rated",
		"upcoming": "upcoming",
	}
	endpoint := endpoints[CLI.Type]
	apiKey := os.Getenv("TMDB_API_KEY")
	if apiKey == "" {
		pterm.Error.Println("TMDB_API_KEY environment variable is not set.")
		pterm.Info.Println("Please set it in your .env file or export it directly.")
		os.Exit(1)
	}
	spinner, _ := pterm.DefaultSpinner.Start(fmt.Sprintf("Fetching %s movies from TMDB...", CLI.Type))
	response, err := api.FetchMovies(endpoint, apiKey)
	if err != nil {
		spinner.Fail("Failed to fetch movies: ", err)
		os.Exit(1)
	}
	spinner.Success("Successfully fetched movies!")
	fmt.Println()
	tableData := pterm.TableData{{"Rank", "Title", "Rating", "Release Date"}}
	for i, movie := range response.Results {
		if i >= 10 {
			break
		}
		ratingStr := fmt.Sprintf("%.1f", movie.VoteAverage)
		if movie.VoteAverage >= 8.0 {
			ratingStr = pterm.Green(ratingStr)
		} else if movie.VoteAverage >= 6.0 {
			ratingStr = pterm.Yellow(ratingStr)
		} else {
			ratingStr = pterm.Red(ratingStr)
		}
		tableData = append(tableData, []string{fmt.Sprintf("%d", i+1), pterm.LightCyan(movie.Title), ratingStr, movie.ReleaseDate})
	}
	err = pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).Render()
	if err != nil {
		pterm.Error.Printf("Failed to render table: %v\n", err)
	}
}
