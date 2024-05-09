package strategies

type Strategy interface {
	Evaluate(pawns, opponentPawns map[[2]int]bool, length, currentPlayer, mainPlayer int) float64
	ChangeCoefficients() bool
}
type DistanceStrategy struct{}

type ReverseDistanceStrategy struct{}

type CrowdedStrategy struct{}

type MainStrategy struct {
	CoefficientPlayerState   float64
	CoefficientOpponentState float64
	CoefficientPlayerCrowded float64
	CoefficientManhattanSum  float64
}

type AdaptiveStrategy struct {
	CoefficientPlayerState   float64
	CoefficientOpponentState float64
	CoefficientPlayerCrowded float64
	CoefficientManhattanSum  float64
	Level                    int
}

func (thisStrategy *DistanceStrategy) Evaluate(playerPawns, opponentPawns map[[2]int]bool, length, currentPlayer, mainPlayer int) float64 {
	return DetermineBoardStateByDistance(playerPawns, opponentPawns, length, currentPlayer, mainPlayer)
}

func (thisStrategy *ReverseDistanceStrategy) Evaluate(playerPawns, opponentPawns map[[2]int]bool, length, currentPlayer, mainPlayer int) float64 {
	if currentPlayer == mainPlayer {
		return -0.75*countManhattanDistanceSum(opponentPawns, length, mainPlayer) +
			0.25*DetermineBoardStateByDistance(playerPawns, opponentPawns, length, currentPlayer, mainPlayer)
	}
	return -0.75*countManhattanDistanceSum(playerPawns, length, currentPlayer) +
		0.25*DetermineBoardStateByDistance(playerPawns, opponentPawns, length, currentPlayer, mainPlayer)
}

func (thisStrategy *CrowdedStrategy) Evaluate(playerPawns, opponentPawns map[[2]int]bool, length, currentPlayer, mainPlayer int) float64 {
	return 0.25*crowdedStrategy(playerPawns, opponentPawns, length, currentPlayer, mainPlayer) +
		0.75*DetermineBoardStateByDistance(playerPawns, opponentPawns, length, currentPlayer, mainPlayer)
}

func (thisStrategy *MainStrategy) Evaluate(playerPawns, opponentPawns map[[2]int]bool, length, currentPlayer, mainPlayer int) float64 {
	return thisStrategy.determineBoardStateByCoefficent(playerPawns, opponentPawns, length, currentPlayer, mainPlayer)
}

func (thisStrategy *AdaptiveStrategy) Evaluate(playerPawns, opponentPawns map[[2]int]bool, length, currentPlayer, mainPlayer int) float64 {
	return thisStrategy.determineBoardStateByCoefficent(playerPawns, opponentPawns, length, currentPlayer, mainPlayer)
}
