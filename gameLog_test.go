package elo

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
	"time"
)

func TestNewGameLog_CreatesAGameLogWithZeroEntries(t *testing.T) {
	actual := NewGameLog()
	expected := GameLog{
		Entries: make([]GameLogEntry, 0),
	}

	assert.Equal(t, expected, actual, "GameLog should have empty slice of GameLogEntries")
}

func TestGameLog_CanBeSorted(t *testing.T) {
	gl := NewGameLog()
	first, _ := time.Parse("_2 Jan 2006 15:04:07", "1 Jan 2016 11:18:12")
	second, _ := time.Parse("_2 Jan 2006 15:04:07", "1 Jan 2016 17:58:40")
	third, _ := time.Parse("_2 Jan 2006 15:04:07", "5 Oct 2016 06:08:10")
	gl.Entries = []GameLogEntry{
		{
			Created: second,
			Winner:  "bruce",
			Loser:   "diana",
		},
		{
			Created: first,
			Winner:  "clark",
			Loser:   "diana",
		},
		{
			Created: third,
			Winner:  "clark",
			Loser:   "bruce",
		},
	}

	sort.Sort(gl)
	assert.Equal(t, first, gl.Entries[0].Created, "Wrong first item")
	assert.Equal(t, second, gl.Entries[1].Created, "Wrong second item")
	assert.Equal(t, third, gl.Entries[2].Created, "Wrong third item")
}
