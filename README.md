# Movie Database CLI 🎬

A beautiful and fast Command Line Interface (CLI) tool built in Go to fetch and display movie data from The Movie Database (TMDB) API. 

Built as a solution for the [roadmap.sh TMDB CLI Project](https://roadmap.sh/projects/tmdb-cli).

---

## Features
* **Live API Integration:** Fetches real-time movie data from TMDB.
* **Beautiful Terminal UI:** Uses `pterm` for loading spinners, boxed tables, and color-coded ratings.
* **Robust Argument Parsing:** Uses `kong` to automatically validate inputs and generate help menus.
* **Secure:** Uses `godotenv` to safely load API keys from a `.env` file.

---

## Prerequisites
1. You must have [Go](https://go.dev/) installed.
2. You need a free API Key from [TMDB](https://www.themoviedb.org/). (Ensure you use the **v4 API Read Access Token**).

---

## Installation & Setup
1. **Clone the repository:**
   ```bash
   git clone https://github.com/lucadani7/MovieDatabaseCLI.git
   cd MovieDatabaseCLI
   ```
2. **Install dependencies:**
   ```bash
   go mod tidy
   ```
3. **Set up your environment variables, creating a `.env` file in the root directory and adding your v4 API token:**
   ```bash
   echo "TMDB_API_KEY=your_v4_read_access_token_here" >> .env
   ```
4. **Build the executable:**
   ```bash
   go build -o tmdb
   ```
---

## Usage
Run the compiled binary and pass the `--type` (or `-t`) flag to specify which movies you want to see.
```bash
# See currently playing movies
./tmdb --type playing

# See popular movies
./tmdb --type popular

# See top-rated movies
./tmdb --type top

# See upcoming movies
./tmdb --type upcoming
```
---

## Help Menu
If you ever forget the commands, just use the built-in help flag:
```bash
./tmdb --help
```
---

## License
This project is licensed under the Apache-2.0 License.

