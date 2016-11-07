package elo

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAddGameToH2HRecords(t *testing.T) {
	recs := make([]H2HRecord, 0)
	recs = AddGameToH2HRecords("tony", recs, GameLogEntry{
		Winner: "steve",
		Loser: "tony",
	})
	recs = AddGameToH2HRecords("tony", recs, GameLogEntry{
		Winner: "steve",
		Loser: "tony",
	})

	assert.Equal(t, "steve", recs[0].Opponent)
	assert.Equal(t, 0, recs[0].Won)
	assert.Equal(t, 2, recs[0].Lost)
}
