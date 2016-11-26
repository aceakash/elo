package main

import (
	"fmt"
	"github.com/aceakash/elo"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"encoding/json"
	"time"
)

type Response struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

func main() {
	table := loadTableFromJsonStore()
	var writeLock sync.Mutex
	r := http.NewServeMux()

	RespondToPoolCommands := func(w http.ResponseWriter, r *http.Request) {
		text := strings.ToLower(r.URL.Query().Get("text"))
		user := strings.ToLower(r.URL.Query().Get("user_name"))

		fmt.Println("Text", text)
		commands := strings.Split(text, " ")
		if len(commands) < 1 {
			fmt.Fprint(w, "Not sure what you want to know. Valid commands are: h2h, help")
		}
		switch commands[0] {
		case "help":
			usage(w)
		case "ratings":
			table = loadTableFromJsonStore()
			fmt.Fprint(w, "```\n")
			fmt.Fprint(w, "#  |                           |      |     |    \n")
			for i, player := range table.GetPlayersSortedByRating() {
				fmt.Fprintf(w, "%2d | %25s | %4d | %3d | %3d \n", i + 1, player.Name, player.Rating, player.Played, player.Won)
			}
			fmt.Fprint(w, "```")
		case "log":
			table = loadTableFromJsonStore()
			fmt.Fprint(w, "```\n")
			var logsToSkip int
			if len(table.GameLog.Entries) > 50 {
				logsToSkip = len(table.GameLog.Entries) - 50
			}
			for i, gle := range table.GameLog.Entries {
				if i < logsToSkip {
					continue
				}
				created := gle.Created.Format("02 Jan")
				fmt.Fprintf(w, "[%s] [%s] %17s (%d -> %d)   defeated %17s (%d -> %d) - added by %s\n", gle.Id, created, gle.Winner, gle.WinnerChange.Before, gle.WinnerChange.After, gle.Loser, gle.LoserChange.Before, gle.LoserChange.After, gle.AddedBy)
			}
			fmt.Fprint(w, "```")
		case "h2h":
			table = loadTableFromJsonStore()
			if len(commands) < 2 {
				usage(w)
				return
			}
			player := commands[1]
			normalisedPlayer := removeAtPrefix(player)
			playerH2HRecords, err := table.HeadToHeadAll(normalisedPlayer)
			if err != nil {
				if err == elo.PlayerDoesNotExist {
					fmt.Fprintf(w, "%s does not seem to be registered", player)
					return
				}
				fmt.Fprint(w, "Sorry, an error occurred processing that ... please contact the admin")
				return
			}
			printH2H(w, player, playerH2HRecords)
		case "result":
			writeLock.Lock()
			defer writeLock.Unlock()
			table = loadTableFromJsonStore()
			if len(commands) < 3 {
				usage(w)
				return
			}
			winner := removeAtPrefix(commands[1])
			loser := removeAtPrefix(commands[2])
			if table.GameLog.HavePlayedOnTheDay(winner, loser, time.Now()) {
				fmt.Fprint(w, "Sorry, those players have already played today! Look at `/pool log`")
				return
			}
			err := table.AddResult(winner, loser, user)
			if err != nil {
				if err == elo.PlayerDoesNotExist {
					fmt.Fprint(w, "Are you sure both those players are registered? Check the ratings table...")
					return
				}
				fmt.Fprint(w, "Sorry, an error occurred processing that ... please contact the admin")
				return
			}
			err = saveTableToJsonStore(table)
			if err != nil {
				fmt.Println("Error writing table to file", err)
				fmt.Fprint(w, "Sorry there was an error - please contact the admin")
				return
			}
			addedGameEntry := table.GameLog.Entries[len(table.GameLog.Entries)-1]
			msg := fmt.Sprintf("Result added by %s: %s defeated %s. Game id is %s", addedGameEntry.AddedBy, addedGameEntry.Winner, addedGameEntry.Loser, addedGameEntry.Id)
			resp := Response{
				Text: msg,
				ResponseType: "in_channel",
			}
			respJson, _ := json.Marshal(resp)
			fmt.Println(string(respJson))
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(respJson))
		default:
			usage(w)
		}
	}

	// Only log requests to our admin dashboard to stdout
	r.Handle("/slack", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(RespondToPoolCommands)))

	port := os.Getenv("PORT")

	fmt.Println("Starting on port ", port)
	err := http.ListenAndServe(":"+port, handlers.CompressHandler(r))
	if err != nil {
		log.Fatal("Error while starting the server", err)
	}
}
func printH2H(w http.ResponseWriter, player string, h2HRecords map[string]*elo.H2HRecord) {
	fmt.Fprintln(w, "```")
	fmt.Fprintf(w, "H2H for %s:\n", player)
	for opponent, h2hr := range h2HRecords {
		fmt.Fprintf(w, "%d - %d  vs %s\n", h2hr.Won, h2hr.Lost, opponent)
	}
	fmt.Fprintln(w, "```")
}

func removeAtPrefix(slackUserName string) string {
	if slackUserName[0] == '@' {
		return slackUserName[1:]
	}
	return slackUserName
}
func usage(w http.ResponseWriter) {
	fmt.Fprintln(w, "Valid commands are:")
	fmt.Fprintln(w, "help: show this help message")
	fmt.Fprintln(w, "ratings: see the ratings table")
	fmt.Fprintln(w, "h2h <player_name>: see player's head-to-head stats vs everyone else")
	fmt.Fprintln(w, "log: see all the games played so far")
	fmt.Fprintln(w, "result <winner> <loser>: add a result")
}

func loadTableFromJsonStore() elo.Table {
	store := elo.JsonFileTableStore{
		Filepath: "eloTable.json",
	}
	table, err := store.Load()
	if err != nil {
		panic(err)
	}
	return table
}

func saveTableToJsonStore(table elo.Table) error {
	store := elo.JsonFileTableStore{
		Filepath: "eloTable.json",
	}
	return store.Save(table)
}

