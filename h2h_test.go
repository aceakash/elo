package elo

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAddGameToH2HRecords_ForSinglePlayer(t *testing.T) {
	// arrange
	recs := make([]H2HRecord, 0)

	// act
	recs = registerGameForPlayerH2H("tony", recs, GameLogEntry{
		Winner: "steve",
		Loser: "tony",
	})
	recs = registerGameForPlayerH2H("tony", recs, GameLogEntry{
		Winner: "steve",
		Loser: "tony",
	})
	recs = registerGameForPlayerH2H("tony", recs, GameLogEntry{
		Winner: "tony",
		Loser: "steve",
	})
	recs = registerGameForPlayerH2H("tony", recs, GameLogEntry{
		Winner: "steve",
		Loser: "tony",
	})

	// assert
	expected := []H2HRecord{
		{
			Opponent: "steve",
			Won: 1,
			Lost: 3,
		},
	}
	assert.Equal(t, expected, recs)
}

func TestTable_HeadToHeadAll_NonExistentPlayer(t *testing.T) {
	// arrange
	table := NewTable(32, 2000)
	table.Register("peter")
	table.Register("natasha")

	// act
	_, err := table.HeadToHeadAll("gruffalo")

	// assert
	assert.Equal(t, PlayerDoesNotExist, err)
}
