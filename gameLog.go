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
		//fmt.Println(game)
		gy, gm, gd := game.Created.Date()
		//fmt.Println("game", gy, gm, gd)
		ny, nm, nd := day.Date()
		//fmt.Println("param", gy, gm, gd)
		if !(gd == nd && gm == nm && gy == ny) {
			continue
		}
		if game.Winner != player1 && game.Winner != player2 {
			continue
		}
		if game.Loser != player1 && game.Loser != player2 {
			continue
		}
		return true
	}
	return false
}

