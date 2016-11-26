package elo

import (
	"testing"

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

func TestTable_AddResult_UpdatesGameCounters(t *testing.T) {
	table := NewTable(32, 2000)
	table.Register("bruce")
	table.Register("clark")
	table.AddResult("bruce", "clark", "")
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
	err := table.AddResult("barry", "bruce", "")
	assert.Equal(t, PlayerDoesNotExist, err, "Did not get expected error")
	err2 := table.AddResult("bruce", "judith", "")
	assert.Equal(t, PlayerDoesNotExist, err2, "Did not get expected error")
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

func TestTable_RecalculateRatingsFromLog_Works(t *testing.T) {
	wrongTable := NewTable(32, 2000)
	wrongTable.Register("bruce")
	wrongTable.Register("clark")
	wrongTable.Register("diana")
	wrongTable.AddResult("bruce", "clark", "")
	wrongTable.AddResult("bruce", "diana", "")
	wrongTable.AddResult("bruce", "diana", "") // double entry
	wrongTable.AddResult("diana", "clark", "")

	accurateTable := NewTable(32, 2000)
	accurateTable.Register("bruce")
	accurateTable.Register("clark")
	accurateTable.Register("diana")
	accurateTable.AddResult("bruce", "clark", "")
	accurateTable.AddResult("bruce", "diana", "")
	accurateTable.AddResult("diana", "clark", "")


	wrongTable.GameLog.Entries = append(wrongTable.GameLog.Entries[0:2], wrongTable.GameLog.Entries[3:]...)
	err := wrongTable.RecalculateRatingsFromLog()

	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, accurateTable.Players["bruce"].Rating, wrongTable.Players["bruce"].Rating, "Bruce has wrong rating")
	assert.Equal(t, accurateTable.Players["clark"].Rating, wrongTable.Players["clark"].Rating, "Clark has wrong rating")
	assert.Equal(t, accurateTable.Players["diana"].Rating, wrongTable.Players["diana"].Rating, "Diana has wrong rating")
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

func TestTable_HeadToHeadAll(t *testing.T) {
	table := NewTable(32, 2000)
	table.Register("steve")
	table.Register("tony")
	table.AddResult("steve", "tony", "")
	table.AddResult("steve", "tony", "")
	table.AddResult("steve", "tony", "")
	table.AddResult("tony", "steve", "")

	recs, err := table.HeadToHeadAll("steve")
	assert.Nil(t, err)
	assert.NotNil(t, recs["tony"], "Must have a H2H record against tony")
	assert.Equal(t, 3, recs["tony"].Won)
	assert.Equal(t, 1, recs["tony"].Lost)
}

func TestTable_GetPlayersSortedByRating(t *testing.T) {
	// arrange
	table := NewTable(32, 2000)
	table.Register("charles")
	table.Register("alice")
	table.Register("bob")

	table.AddResult("alice", "bob", "")
	table.AddResult("alice", "bob", "")
	table.AddResult("alice", "charles", "")
	table.AddResult("bob", "charles", "")
	table.AddResult("bob", "charles", "")
	table.AddResult("charles", "bob", "")

	// act
	players := table.GetPlayersSortedByRating()

	// assert
	assert.Equal(t, "alice", players[0].Name)
	assert.Equal(t, "bob", players[1].Name)
	assert.Equal(t, "charles", players[2].Name)
}
