package elo

import (
	"fmt"
	"math"
)

// Table holds the ratings of all registered players.
type Table struct {
	constantFactor int
	players        map[string]Player
	initialRating  int
}

// NewTable creates a new table.
func NewTable(constantFactor int, initialRating int) Table {
	return Table{
		constantFactor: constantFactor,
		initialRating:  initialRating,
		players:        make(map[string]Player),
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
func (table *Table) AddResult(winnerName, loserName string) {
	winner, _ := table.players[winnerName]
	loser, _ := table.players[loserName]

	winner.played++
	winner.won++

	loser.played++
	loser.lost++

	table.players[winnerName] = winner
	table.players[loserName] = loser

}

func getNewEloRatings(oldRatingForPlayer1 int, oldRatingForPlayer2 int, k int) (int, int) {
	eloDiff := oldRatingForPlayer2 - oldRatingForPlayer1
	perc := 1 / (1 + math.Pow(10.0, float64(eloDiff)/400))
	win := int(math.Floor(float64(k)*(1-perc) + 0.5))
	fmt.Println(oldRatingForPlayer1, oldRatingForPlayer2, win)
	return oldRatingForPlayer1 + win, oldRatingForPlayer2 - win
}
