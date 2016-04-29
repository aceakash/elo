package elo

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
