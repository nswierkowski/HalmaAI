package strategies

func DetermineBoardStateByDistance(playerPawns, opponentPawns map[[2]int]bool, length, currentPlayer, mainPlayer int) float64 {
	if currentPlayer == mainPlayer {
		return countOpponentsResults(opponentPawns, currentPlayer)
	}
	return countPlayerResults(playerPawns, currentPlayer)
}

func (thisStrategy *DistanceStrategy) ChangeCoefficients() bool {
	return false
}
