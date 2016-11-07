package elo

import "time"

type GameLog struct {
	Entries []GameLogEntry `json:entries`
}

type GameLogEntry struct {
	Created      time.Time    `json:created`
	Winner       string       `json:winner`
	Loser        string       `json:loser`
	Notes        string       `json:notes`
	WinnerChange RatingChange `json:winnerChange`
	LoserChange  RatingChange `json:loserChange`
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
