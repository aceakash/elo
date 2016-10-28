package elo

type Player struct {
	Name   string
	Rating int
	Played int
	Won    int
	Lost   int
}

type Players []Player

func (players Players) Len() int {
	return len(players)
}

func (players Players) Less(i, j int) bool {
	return players[i].Rating > players[j].Rating
}

func (players Players) Swap(i, j int) {
	players[i], players[j] = players[j], players[i]
}
