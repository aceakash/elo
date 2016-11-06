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
		user := strings.ToLower(r.URL.Query().Get("user_name"))

		commands := strings.Split(text, " ")
		if len(commands) < 1 {
			fmt.Fprint(w, "Not sure what you want to know. Valid commands are: h2h, help")
		}
		switch commands[0] {
		case "help":
			usage(w)
		//case "ratings":
		//	for _, player := range table.GetPlayersSortedByRating() {
		//		fmt.Fprintf(w, "\n%25s (%d) - Played %2d, Won %2d, Lost %2d", player.Name, player.Rating, player.Played, player.Won, player.Lost)
		//	}
		//case "gamelog":
		//	for _, gle := range table.GameLog.Entries {
		//		created := gle.Created.Format("_2 Jan 2006")
		//		fmt.Fprintf(w, "[%s] %17s (%d -> %d)   defeated %17s (%d -> %d)\n", created, gle.Winner, gle.WinnerChange.Before, gle.WinnerChange.After, gle.Loser, gle.LoserChange.Before, gle.LoserChange.After)
		//	}
		case "h2h":
			if len(commands) < 2 {
				usage(w)
				return
			}
			against := commands[1]
			normalisedAgainst := removeAtPrefix(against)
			userPts, againstPts := table.HeadToHead(user, normalisedAgainst)
			fmt.Fprintf(w, "You are %d - %d against %s", userPts, againstPts, against)
		default:
			fmt.Fprintf(w, "[%s] %s", time.Now().Format(time.RFC3339), text)
		}
	}

	// Only log requests to our admin dashboard to stdout
	r.Handle("/slack", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(RespondToPoolCommands)))

	port := os.Getenv("PORT")

	fmt.Println("Starting on port ", port)
	err = http.ListenAndServe(":"+port, handlers.CompressHandler(r))
	if err != nil {
		log.Fatal("Error while starting the server", err)
	}
}
func removeAtPrefix(slackUserName string) string {
	return slackUserName[1:]
}
func usage(w http.ResponseWriter) {
	fmt.Fprint(w, "Valid commands are:\n")
	fmt.Fprint(w, "help: show this help message\n")
	fmt.Fprint(w, "h2h <@another_player>: see your head-to-head stats vs another player")
}
