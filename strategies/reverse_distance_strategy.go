package strategies

func reverseDetermineBoardStateByDistance(playerPawns, opponentPawns map[[2]int]bool, length, currentPlayer, mainPlayer int) float64 {
	if currentPlayer == mainPlayer {
		return countPlayerResults(playerPawns, currentPlayer)
	}
	return countOpponentsResults(opponentPawns, currentPlayer)
}

func (thisStrategy *ReverseDistanceStrategy) ChangeCoefficients() bool {
	return false
}
