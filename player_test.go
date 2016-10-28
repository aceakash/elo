package elo

import (
	"testing"
	"sort"
	"github.com/stretchr/testify/assert"
)

func TestPlayersCanBeSortedDescByRating(t *testing.T) {
	playersSlice := []Player{
		{Rating: 3000, Name: "Charlie", Lost: 2, Played: 3, Won: 1,},
		{Rating: 2000, Name: "Bob", Lost: 2, Played: 3, Won: 1,},
		{Rating: 4000, Name: "Dan", Lost: 2, Played: 3, Won: 1,},
		{Rating: 1000, Name: "Alice", Lost: 2, Played: 3, Won: 1,},
	}
	sort.Sort(Players(playersSlice))

	assert.Equal(t, "Dan", playersSlice[0].Name, "Expected Dan to be first")
	assert.Equal(t, "Charlie", playersSlice[1].Name, "Expected Charlie to be first")
	assert.Equal(t, "Bob", playersSlice[2].Name, "Expected Bob to be first")
	assert.Equal(t, "Alice", playersSlice[3].Name, "Expected Alice to be first")
}
