package elo

import (
	"time"
)

type GameLog struct {
	Entries []GameLogEntry `json:entries`
}

type GameLogEntry struct {
	Id           string       `json:id`
	Created      time.Time    `json:created`
	Winner       string       `json:winner`
	Loser        string       `json:loser`
	Notes        string       `json:notes`
	WinnerChange RatingChange `json:winnerChange`
	LoserChange  RatingChange `json:loserChange`
	AddedBy      string       `json:addedBy`
}

type RatingChange struct {
	Before int `json:before`
	After  int `json:after`
}

func NewGameLog() GameLog {
	return GameLog{
		Entries: make([]GameLogEntry, 0),
	}
}

func (gl GameLog) Len() int {
	return len(gl.Entries)
}

func (gl GameLog) Less(i, j int) bool {
	return gl.Entries[i].Created.Unix() < gl.Entries[j].Created.Unix()
}

func (gl GameLog) Swap(i, j int) {
	gl.Entries[i], gl.Entries[j] = gl.Entries[j], gl.Entries[i]
}

func (gl GameLog) HavePlayedOnTheDay(player1 string, player2 string, day time.Time) bool {
	for _, game := range gl.Entries {
		sameDay := isSameDay(day, game.Created)
		player1Found := player1 == game.Winner || player1 == game.Loser
		player2Found := player2 == game.Winner || player2 == game.Loser
		if sameDay && player1Found && player2Found {
			return true
		}
	}
	return false
}

func isSameDay(time1, time2 time.Time) bool {
	gy, gm, gd := time1.Date()
	ny, nm, nd := time2.Date()
	return gd == nd && gm == nm && gy == ny
}

// GetOpponent gets the opponent name for a player from a game.
// Returns "" if player did not play in the game.
func GetOpponent(game GameLogEntry, player string) string {
	switch {
	case player == game.Winner:
		return game.Loser
	case player == game.Loser:
		return game.Winner
	default:
		return ""
	}
}

