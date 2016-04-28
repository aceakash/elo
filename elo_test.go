package elo

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type Player struct {
	name string
	rating int
	played int
	won int
	lost int
}

type EloTable struct {
	ConstantFactor int
	Players map[string]Player
	InitialRating int
}

func NewEloTable(constantFactor int, initialRating int) EloTable {
	return EloTable{
		ConstantFactor: constantFactor,
		InitialRating: initialRating,
		Players: make(map[string]Player),
	}
}

func TestNewEloTableSetsCorrectValues(t *testing.T) {
	table := NewEloTable(32, 2000)
	assert.Equal(t, 32, table.ConstantFactor, "ConstantFactor is wrong");
	assert.Equal(t, 2000, table.InitialRating, "InitialRating is wrong");
	assert.Empty(t, table.Players, "Players map not empty");
}
