package elo

import "time"

type GameLog struct {
	Entries []GameLogEntry `json:entries`
}

type GameLogEntry struct {
	Created time.Time `json:created`
	Winner string `json:winner`
	Loser string `json:loser`
	Notes string `json:notes`
}
