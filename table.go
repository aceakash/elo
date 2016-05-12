package elo

import "time"

// Table holds the ratings of all registered players.
type Table struct {
	constantFactor int
	initialRating  int
	players        map[string]Player
	Matches        []Match
}

// NewTable creates a new table.
func NewTable(constantFactor int, initialRating int) Table {
	return Table{
		constantFactor: constantFactor,
		initialRating:  initialRating,
		players:        make(map[string]Player),
		Matches:        make([]Match, 0, 100),
	}
}

// Register adds a new player to the table.
func (table *Table) Register(playerName string) {
	table.players[playerName] = Player{
		rating: table.initialRating,
		name:   playerName,
	}
}

// AddResult adds the result of a match to the table.
func (table *Table) AddResult(winner, loser string) {
	winningPlayer, _ := table.players[winner]
	losingPlayer, _ := table.players[loser]
	winningPlayer.played = 1
	winningPlayer.lost = 0
	winningPlayer.won = 1
	losingPlayer.played = 1
	losingPlayer.lost = 1
	losingPlayer.won = 0
	table.players[winner] = winningPlayer
	table.players[loser] = losingPlayer
	match := Match{
		WinnerName: winner,
		LoserName:  loser,
		RecordedOn: time.Now(),
		Notes:      "was a close match!",
	}
	table.Matches = append(table.Matches, match)
}
