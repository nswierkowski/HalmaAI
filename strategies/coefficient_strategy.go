package strategies

func (thisStrategy *MainStrategy) determineBoardStateByCoefficent(playerPawns, opponentPawns map[[2]int]bool, length, currentPlayer, mainPlayer int) float64 {
	return thisStrategy.CoefficientPlayerState*DetermineBoardStateByDistance(playerPawns, opponentPawns, length, currentPlayer, mainPlayer) +
		thisStrategy.CoefficientOpponentState*reverseDetermineBoardStateByDistance(playerPawns, opponentPawns, length, currentPlayer, mainPlayer) +
		thisStrategy.CoefficientPlayerCrowded*crowdedStrategy(playerPawns, opponentPawns, length, currentPlayer, mainPlayer) +
		thisStrategy.CoefficientManhattanSum*countManhattanDistanceSum(playerPawns, length, currentPlayer)
}

func (thisStrategy *MainStrategy) ChangeCoefficients() bool {
	return false
}
