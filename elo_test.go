package elo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEloTableSetsValues(t *testing.T) {
	table := NewTable(32, 2000)
	assert.Equal(t, 32, table.constantFactor, "ConstantFactor is wrong")
	assert.Equal(t, 2000, table.initialRating, "InitialRating is wrong")
	assert.Empty(t, table.players, "Players map not empty")
}

func TestRegisterPlayer(t *testing.T) {
	table := NewTable(32, 2000)
	table.Register("bruce")
	assert.Equal(t, 2000, table.players["bruce"].rating, "initial rating wrong")
	assert.Equal(t, "bruce", table.players["bruce"].name, "name wrong")
	assert.Equal(t, 0, table.players["bruce"].played, "played should be 0")
	assert.Equal(t, 0, table.players["bruce"].won, "won should be 0")
	assert.Equal(t, 0, table.players["bruce"].lost, "lost should be 0")
}

func TestAddResult(t *testing.T) {
	table := NewTable(32, 2000)
	table.Register("bruce")
	table.Register("clark")
	bruce := table.players["bruce"]
	fmt.Println(bruce)
	table.AddResult("bruce", "clark")
	assert.Equal(t, 1, table.players["bruce"].played, "bruce should have played 1 game")
	assert.Equal(t, 1, table.players["bruce"].won, "bruce should have won 1 game")
	assert.Equal(t, 0, table.players["bruce"].lost, "bruce should have lost 0 games")
	assert.Equal(t, 1, table.players["clark"].played, "clark should have played 1 game")
	assert.Equal(t, 0, table.players["clark"].won, "clark should have won 0 games")
	assert.Equal(t, 1, table.players["clark"].lost, "clark should have lost 1 game")
	assert.Equal(t, 1, len(table.matches))
	assert.Equal(t, "bruce", table.matches[0].winnerName)
	assert.Equal(t, "clark", table.matches[0].loserName)
}
