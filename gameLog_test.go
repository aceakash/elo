package elo

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewGameLog_CreatesAGameLogWithZeroEntries(t *testing.T) {
	actual := NewGameLog()
	expected := GameLog{
		Entries: make([]GameLogEntry, 0),
	}

	assert.Equal(t, expected, actual, "GameLog should have empty slice of GameLogEntries")
}

