package elo

import "math"

// CalculateRating calculates new ELO ratings.
// based on https://metinmediamath.wordpress.com/2013/11/27/how-to-calculate-the-elo-rating-including-example/
func CalculateRating(winnerOldRating int, loserOldRating int, constantFactor int) (int, int) {
	winnerTransformedRating := transformRating(winnerOldRating)
	loserTransformedRating := transformRating(loserOldRating)
	expectedWinnerScore := getExpectedScore(winnerTransformedRating, loserTransformedRating)
	expectedLoserScore := getExpectedScore(loserTransformedRating, winnerTransformedRating)

	winnerNewRating := float64(winnerOldRating) + float64(constantFactor)*(1-expectedWinnerScore)
	loserNewRating := float64(loserOldRating) + float64(constantFactor)*(0-expectedLoserScore)
	return round(winnerNewRating), round(loserNewRating)
}

func transformRating(rating int) float64 {
	power := float64(rating) / 400
	return math.Pow(10, power)
}

func getExpectedScore(transformedRating float64, opponentsTransformedRating float64) float64 {
	return transformedRating / (transformedRating + opponentsTransformedRating)
}

func round(val float64) int {
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
}
