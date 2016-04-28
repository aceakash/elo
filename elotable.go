package elo

type EloTable struct {
	constantFactor int
	players        map[string]Player
	initialRating  int
}

func NewEloTable(constantFactor int, initialRating int) EloTable {
	return EloTable{
		constantFactor: constantFactor,
		initialRating:  initialRating,
		players:        make(map[string]Player),
	}
}
