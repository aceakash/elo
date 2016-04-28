package elo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEloTableSetsValues(t *testing.T) {
	table := NewEloTable(32, 2000)
	assert.Equal(t, 32, table.constantFactor, "ConstantFactor is wrong")
	assert.Equal(t, 2000, table.initialRating, "InitialRating is wrong")
	assert.Empty(t, table.players, "Players map not empty")
}

func TestRegisterPlayer(t *testing.T) {
	table := NewEloTable(32, 2000)
	table.Register("bruce")
	assert.Equal(t, 2000, table.players["bruce"].rating, "initial rating wrong")
	assert.Equal(t, "bruce", table.players["bruce"].name, "name wrong")
	assert.Equal(t, 0, table.players["bruce"].played, "played should be 0")
	assert.Equal(t, 0, table.players["bruce"].won, "won should be 0")
	assert.Equal(t, 0, table.players["bruce"].lost, "lost should be 0")
}

func (table *EloTable) Register(playerName string) {
	table.players[playerName] = Player{
		rating: table.initialRating,
		name: playerName,
	}
}
//func TestAddResult(t *testing.T) {
//	table := NewEloTable(32, 2000)
//	table.AddResult("bruce", "clark")
//	assert.Equal(t, table[""])
//}
