package elo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateRating(t *testing.T) {
	// validated against http://www.3dkingdoms.com/chess/elo.htm
	testData := []struct {
		winnerRating         int
		loserRating          int
		constantFactor       int
		expectedWinnerRating int
		expectedLoserRating  int
	}{
		{1000, 1000, 24, 1012, 988},
		{1000, 1000, 32, 1016, 984},
		{2100, 2400, 32, 2127, 2373},
		{1000, 6000, 10, 1010, 5990},
	}

	for _, td := range testData {
		winnerNewRating, loserNewRating := CalculateRating(td.winnerRating, td.loserRating, td.constantFactor)
		assert.Equal(t, td.expectedWinnerRating, winnerNewRating, "Wrong new rating for winner")
		assert.Equal(t, td.expectedLoserRating, loserNewRating, "Wrong new rating for loser")
	}
}
