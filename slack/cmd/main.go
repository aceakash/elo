package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"github.com/aceakash/elo"
)

func main() {
	store := elo.JsonFileTableStore{
		Filepath: "eloTable.json",
	}
	table, err := store.Load()
	if err != nil {
		panic(err)
	}

	r := http.NewServeMux()

	RespondToPoolCommands := func(w http.ResponseWriter, r *http.Request) {
		text := strings.ToLower(r.URL.Query().Get("text"))
		switch text {
		case "ratings":
			for _, player := range table.GetPlayersSortedByRating() {
				fmt.Fprintf(w, "\n%25s (%d) - Played %2d, Won %2d, Lost %2d", player.Name, player.Rating, player.Played, player.Won, player.Lost)
			}
		case "gamelog":
			for _, gle := range table.GameLog.Entries {
				created := gle.Created.Format("_2 Jan 2006")
				fmt.Fprintf(w, "[%s] %17s (%d -> %d)   defeated %17s (%d -> %d)\n", created, gle.Winner, gle.WinnerChange.Before, gle.WinnerChange.After, gle.Loser, gle.LoserChange.Before, gle.LoserChange.After)
			}
		default:
			fmt.Fprintf(w, "[%s] %s", time.Now().Format(time.RFC3339), text)
		}
	}

	// Only log requests to our admin dashboard to stdout
	r.Handle("/pool-command", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(RespondToPoolCommands)))

	port := os.Getenv("PORT")

	fmt.Println("Starting on port ", port)
	err = http.ListenAndServe(":"+port, handlers.CompressHandler(r))
	if err != nil {
		log.Fatal("Error while starting the server", err)
	}
}
