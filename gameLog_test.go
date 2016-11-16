package elo

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
	"time"
	"fmt"
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
	first, _ := time.Parse("_2 Jan 2006 15:04:05", "1 Jan 2016 11:18:12")
	second, _ := time.Parse("_2 Jan 2006 15:04:05", "1 Jan 2016 17:58:40")
	third, _ := time.Parse("_2 Jan 2006 15:04:05", "5 Oct 2016 06:08:10")
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

func TestGameLog_HavePlayedOnTheDay(t *testing.T) {
	gl := NewGameLog()
	format := "_2 Jan 2006 15:04:05"
	firstJanAM, err := time.Parse(format, "1 Jan 2016 11:18:12")
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(firstJanAM)
	second, _ := time.Parse(format, "2 Jan 2016 11:58:40")
	third, _ := time.Parse(format, "3 Jan 2016 11:58:40")
	gl.Entries = []GameLogEntry{
		{
			Created: firstJanAM,
			Winner: "alice",
			Loser: "bob",
		},
		{
			Created: second,
			Winner: "alice",
			Loser: "bob",
		},
	}
	firstJanPM, _ := time.Parse(format, "1 Jan 2016 19:00:00")
	actual := gl.HavePlayedOnTheDay("alice", "bob", firstJanPM)
	assert.True(t, actual, "alice and bob HAVE played on this day")

	actual = gl.HavePlayedOnTheDay("bob", "alice", firstJanPM)
	assert.True(t, actual, "alice and bob HAVE played on this day")

	actual = gl.HavePlayedOnTheDay("alice", "charles", second)
	assert.False(t, actual, "alice and charles HAVE NOT played on this day")

	actual = gl.HavePlayedOnTheDay("alice", "charles", firstJanPM)
	assert.False(t, actual, "alice and charles HAVE NOT played on this day")

	actual = gl.HavePlayedOnTheDay("alice", "bob", third)
	assert.False(t, actual, "alice and bob HAVE NOT played on this day")
}
