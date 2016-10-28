package elo

import (
	"strings"
	"time"
)


// Table holds the ratings of all registered players.
type Table struct {
	ConstantFactor int `json:constantFactor`
	Players        map[string]Player `json:players`
	InitialRating  int `json:initialRating`
	GameLog		GameLog	`json:gameLog`
}

// NewTable creates a new table.
func NewTable(constantFactor int, initialRating int) Table {
	return Table{
		ConstantFactor: constantFactor,
		InitialRating:  initialRating,
		Players:        make(map[string]Player),
		GameLog:		NewGameLog(),
	}
}

func NewGameLog() GameLog {
	return GameLog{
		Entries: make([]GameLogEntry, 0),
	}
}

// Register adds a new player to the table.
func (table *Table) Register(playerName string) error {
	playerName = sanitiseName(playerName)
	if _, exists := table.Players[playerName]; exists {
		return PlayerAlreadyExists
	}
	table.Players[playerName] = Player{
		Rating: table.InitialRating,
		Name:   playerName,
	}
	return nil
}
func sanitiseName(name string) string {
	return strings.ToLower(name)
}

// AddResult adds the result of a match to the table.

func (table *Table) AddResult(winner, loser string) error {
	winningPlayer, exists := table.Players[winner]
	if !exists {
		return PlayerDoesNotExist
	}
	losingPlayer, exists := table.Players[loser]
	if !exists {
		return PlayerDoesNotExist
	}
	winningPlayer.Played++
	winningPlayer.Won++
	losingPlayer.Played++
	losingPlayer.Lost++
	w, l := CalculateRating(winningPlayer.Rating, losingPlayer.Rating, table.ConstantFactor)
	winningPlayer.Rating = w
	losingPlayer.Rating = l
	table.Players[winner] = winningPlayer
	table.Players[loser] = losingPlayer
	gle := GameLogEntry{
		Created: time.Now(),
		Winner: winner,
		Loser: loser,
		Notes: "",
	}
	table.GameLog.Entries = append(table.GameLog.Entries, gle)
	return nil
}



