package elo

// Table holds the ratings of all registered players.
type Table struct {
	ConstantFactor int `json:constantFactor`
	Players        map[string]Player `json:players`
	InitialRating  int `json:initialRating`
}

// NewTable creates a new table.
func NewTable(constantFactor int, initialRating int) Table {
	return Table{
		ConstantFactor: constantFactor,
		InitialRating:  initialRating,
		Players:        make(map[string]Player),
	}
}

// Register adds a new player to the table.
func (table *Table) Register(playerName string) {
	table.Players[playerName] = Player{
		Rating: table.InitialRating,
		Name:   playerName,
	}
}

// AddResult adds the result of a match to the table.
func (table *Table) AddResult(winner, loser string) {
	winningPlayer, _ := table.Players[winner]
	losingPlayer, _ := table.Players[loser]
	winningPlayer.Played = 1
	winningPlayer.Lost = 0
	winningPlayer.Won = 1
	losingPlayer.Played = 1
	losingPlayer.Lost = 1
	losingPlayer.Won = 0
	w, l := CalculateRating(winningPlayer.Rating, losingPlayer.Rating, table.ConstantFactor)
	winningPlayer.Rating = w
	losingPlayer.Rating = l
	table.Players[winner] = winningPlayer
	table.Players[loser] = losingPlayer
}



