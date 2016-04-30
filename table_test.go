package elo

import (
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
	assert.Equal(t, 0, table.players["bruce"].lost, "won should be 0")
	assert.Equal(t, 0, table.players["bruce"].lost, "lost should be 0")
}

func TestAddResultUpdatesNewPlayersPlayedWonLostCounts(t *testing.T) {
	table := NewTable(32, 2000)
	table.Register("bruce")
	table.Register("clark")
	table.AddResult("bruce", "clark")
	assert.Equal(t, 1, table.players["bruce"].played, "bruce should have played 1 game")
	assert.Equal(t, 1, table.players["bruce"].won, "bruce should have won 1 game")
	assert.Equal(t, 0, table.players["bruce"].lost, "bruce should have lost 0 games")

	assert.Equal(t, 1, table.players["clark"].played, "clark should have played 1 game")
	assert.Equal(t, 0, table.players["clark"].won, "clark should have won 0 games")
	assert.Equal(t, 1, table.players["clark"].lost, "clark should have lost 1 game")
}

func TestAddResultUpdatesExistingPlayersPlayedWonLostCounts(t *testing.T) {
	table := NewTable(32, 2000)
	table.Register("bruce")
	table.Register("clark")
	bruce := table.players["bruce"]
	bruce.played = 5
	bruce.won = 3
	bruce.lost = 2
	bruce.rating = 2100
	table.players["bruce"] = bruce

	clark := table.players["clark"]
	clark.played = 6
	clark.won = 1
	clark.lost = 5
	clark.rating = 1800
	table.players["clark"] = clark

	// act
	table.AddResult("clark", "bruce")

	assert.Equal(t, 6, table.players["bruce"].played, "bruce should have played 6")
	assert.Equal(t, 3, table.players["bruce"].won, "bruce should have won 3")
	assert.Equal(t, 3, table.players["bruce"].lost, "bruce should have lost 3")
	assert.Equal(t, 7, table.players["clark"].played, "clark should have played 7")
	assert.Equal(t, 2, table.players["clark"].won, "clark should have won 2")
	assert.Equal(t, 5, table.players["clark"].lost, "clark should have lost 5")
}

func TestAddResultUpdatesRating(t *testing.T) {
	table := getStdTableWithTwoPlayers(Player{name: "bruce", rating: 2100}, Player{name: "clark", rating: 1800})

	table.AddResult("bruce", "clark")

	bruce := table.players["bruce"]
	clark := table.players["clark"]
	assert.Equal(t, 2105, bruce.rating, "rating for bruce is wrong")
	assert.Equal(t, 1795, clark.rating, "rating for clark is wrong")
}

func TestGetNewEloRatings(t *testing.T) {
	var tests = []struct {
		old1 int
		old2 int
		k    int
		new1 int
		new2 int
	}{
		{2000, 2000, 32, 2000 + 16, 2000 - 16},
		{2100, 1800, 32, 2100 + 5, 1800 - 5},
		{2050, 2000, 32, 2050 + 14, 2000 - 14},
	}

	for _, tt := range tests {
		new1, new2 := getNewEloRatings(tt.old1, tt.old2, tt.k)
		assert.Equal(t, tt.new1, new1, "new rating for player1 is wrong")
		assert.Equal(t, tt.new2, new2, "new rating for player2 is wrong")
	}

}

func getStdTableWithTwoPlayers(player1 Player, player2 Player) Table {
	table := NewTable(32, 2000)
	table.players[player1.name] = player1
	table.players[player2.name] = player2
	return table
}
