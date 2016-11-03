package elo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewEloTableSetsValues(t *testing.T) {
	table := NewTable(32, 2000)
	assert.Equal(t, 32, table.ConstantFactor, "ConstantFactor is wrong")
	assert.Equal(t, 2000, table.InitialRating, "InitialRating is wrong")
	assert.Empty(t, table.Players, "Players map not empty")
}

func TestRegisterPlayer(t *testing.T) {
	table := NewTable(32, 2000)
	table.Register("bruce")
	assert.Equal(t, 2000, table.Players["bruce"].Rating, "initial rating wrong")
	assert.Equal(t, "bruce", table.Players["bruce"].Name, "name wrong")
	assert.Equal(t, 0, table.Players["bruce"].Played, "played should be 0")
	assert.Equal(t, 0, table.Players["bruce"].Won, "won should be 0")
	assert.Equal(t, 0, table.Players["bruce"].Lost, "lost should be 0")
}

func TestAddResultUpdatesGameCounters(t *testing.T) {
	table := NewTable(32, 2000)
	table.Register("bruce")
	table.Register("clark")
	table.AddResult("bruce", "clark")
	assert.Equal(t, 1, table.Players["bruce"].Played, "bruce should have played 1 game")
	assert.Equal(t, 1, table.Players["bruce"].Won, "bruce should have won 1 game")
	assert.Equal(t, 0, table.Players["bruce"].Lost, "bruce should have lost 0 games")
	assert.Equal(t, 1, table.Players["clark"].Played, "clark should have played 1 game")
	assert.Equal(t, 0, table.Players["clark"].Won, "clark should have won 0 games")
	assert.Equal(t, 1, table.Players["clark"].Lost, "clark should have lost 1 game")
}

func TestTable_AddResult_ReturnsErrorForNonExistentPlayer(t *testing.T) {
	table := NewTable(32, 2000)
	table.Register("bruce")
	table.Register("clark")
	err := table.AddResult("barry", "bruce")
	assert.Equal(t, PlayerDoesNotExist, err, "Did not get expected error")
}

func TestTable_RecalculateRatingsFromLog_ReturnsEmptyTable_ForEmptyGameLog(t *testing.T) {
	table := NewTable(24, 1000)
	err := table.RecalculateRatingsFromLog()
	if err != nil {
		t.Fatal(err)
	}
	assert.Empty(t, table.Players, "There should be no players")
	assert.Empty(t, table.GameLog.Entries, "There should be nothing in the game log")
}

func TestTable_RecalculateRatingsFromLog_ReturnsRightTable_ForSingleEntryGameLog(t *testing.T) {
	table := NewTable(32, 2000)
	firstOfJan, _ := time.Parse("2006-01-02", "2012-01-01")
	table.GameLog.Entries = append(table.GameLog.Entries, GameLogEntry{
		Created: firstOfJan,
		Winner:  "clark",
		Loser:   "bruce",
	})

	err := table.RecalculateRatingsFromLog()

	if err != nil {
		t.Fatal(err)
	}
	expectedBruce := Player{
		Name:   "bruce",
		Played: 1,
		Won:    0,
		Lost:   1,
		Rating: 1984,
	}
	expectedClark := Player{
		Name:   "clark",
		Played: 1,
		Won:    1,
		Lost:   0,
		Rating: 2016,
	}
	assert.Equal(t, expectedBruce, table.Players["bruce"], "Bruce was not as expected")
	assert.Equal(t, expectedClark, table.Players["clark"], "Clark was not as expected")
}
