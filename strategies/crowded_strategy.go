package strategies

func crowdCoefficientPlayer1(pawns map[[2]int]bool) float64 {
	coefficient := 0.0
	for pawn := range pawns {
		if pawns[[2]int{pawn[0] + 1, pawn[1]}] {
			coefficient += 1.0
		}
		if pawns[[2]int{pawn[0], pawn[1] + 1}] {
			coefficient += 1.0
		}
		if pawns[[2]int{pawn[0] + 1, pawn[1] + 1}] {
			coefficient += 1.0
		}
	}
	return coefficient
}

func crowdCoefficientPlayer2(pawns map[[2]int]bool) float64 {
	coefficient := 0.0
	for pawn := range pawns {
		if pawns[[2]int{pawn[0] - 1, pawn[1]}] {
			coefficient += 1.0
		}
		if pawns[[2]int{pawn[0], pawn[1] - 1}] {
			coefficient += 1.0
		}
		if pawns[[2]int{pawn[0] - 1, pawn[1] - 1}] {
			coefficient += 1.0
		}
	}
	return coefficient
}

func crowdedStrategy(playerPawns, opponentPawns map[[2]int]bool, length, currentPlayer, mainPlayer int) float64 {
	if currentPlayer == mainPlayer {
		if currentPlayer == 1 {
			return crowdCoefficientPlayer1(playerPawns)
		}
		return crowdCoefficientPlayer2(playerPawns)
	}

	if currentPlayer == 1 {
		return crowdCoefficientPlayer1(opponentPawns)
	}
	return crowdCoefficientPlayer2(opponentPawns)
}

func (thisStrategy *CrowdedStrategy) ChangeCoefficients() bool {
	return false
}
