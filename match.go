package elo

import (
	"bufio"
	"encoding/gob"
	"os"
	"time"
)

// Match records the outcome of a 2-player match. There needs to be a clear winner.
type Match struct {
	WinnerName string
	LoserName  string
	RecordedOn time.Time
	Notes      string
}

// PersistMatches persists matches to a datastore.
func PersistMatches(matches []Match) {
	f, err := os.Create("matches.gob")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	w := bufio.NewWriter(f)
	enc := gob.NewEncoder(w)
	err = enc.Encode(Match{"akash", "yash", time.Date(2016, 5, 11, 0, 0, 0, 0, time.UTC), ""})
	if err != nil {
		panic(err)
	}
	if err := w.Flush(); err != nil {
		panic(err)
	}
}
