package strategies

import "fmt"

var coefficientsPerMoveMean = map[int][3]float64{
	0:  {1.0167, -0.3, 0.45},
	10: {0.95, -0.3, 0.65},
	20: {0.883, -0.3, 0.383},
	30: {1.083, -0.4, 0.483},
	40: {1.0167, -0.4, 0.25},
	50: {0.95, -0.167, 0.45},
	60: {1.083, -0.23, 0.283},
	70: {1.1167, -0.167, 0.45},
	80: {1.183, -0.23, 0.65},
	90: {0.883, -0.3, 0.283},
}

var coefficientsPerMove = map[int][3]float64{
	0:  {1.15, -0.4, 0.35},
	10: {1.55, -0.2, 0.85},
	20: {1.15, -0.4, 0.35},
	30: {1.15, -0.6, 0.45},
	40: {1.15, -0.4, 0.05},
	50: {1.15, -0.2, 0.65},
	60: {1.25, -0.4, 0.45},
	70: {1.55, -0.1, 0.45},
	80: {1.25, -0.2, 0.65},
	90: {1.15, 0.1, 0.35},
}

var coefficientsPerMoveMedian = map[int][3]float64{
	0:  {1.05, -0.4, 0.35},
	10: {0.75, -0.3, 0.85},
	20: {0.85, -0.4, 0.35},
	30: {1.15, -0.4, 0.35},
	40: {1.05, -0.4, 0.35},
	50: {1.05, -0.2, 0.25},
	60: {1.05, -0.4, 0.25},
	70: {1.05, -0.1, 0.45},
	80: {1.25, -0.2, 0.65},
	90: {0.75, -0.5, 0.25},
}

func (thisStrategy *AdaptiveStrategy) determineBoardStateByCoefficent(playerPawns, opponentPawns map[[2]int]bool, length, currentPlayer, mainPlayer int) float64 {
	return thisStrategy.CoefficientPlayerState*DetermineBoardStateByDistance(playerPawns, opponentPawns, length, currentPlayer, mainPlayer) +
		thisStrategy.CoefficientOpponentState*reverseDetermineBoardStateByDistance(playerPawns, opponentPawns, length, currentPlayer, mainPlayer) +
		thisStrategy.CoefficientPlayerCrowded*crowdedStrategy(playerPawns, opponentPawns, length, currentPlayer, mainPlayer) +
		thisStrategy.CoefficientManhattanSum*countManhattanDistanceSum(playerPawns, length, currentPlayer)
}

func (thisStrategy *AdaptiveStrategy) ChangeCoefficients() bool {
	thisStrategy.Level++

	if thisStrategy.Level/10 > 9 {
		return false
	}

	fmt.Println("##############################################################################################")
	fmt.Println("LEVEL: ", thisStrategy.Level-(thisStrategy.Level%10))
	fmt.Println("##############################################################################################")
	coefficients := coefficientsPerMove[thisStrategy.Level-(thisStrategy.Level%10)]
	thisStrategy.CoefficientPlayerState = coefficients[0]
	thisStrategy.CoefficientOpponentState = coefficients[1]
	thisStrategy.CoefficientPlayerCrowded = coefficients[2]
	return true
}
