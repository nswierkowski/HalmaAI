package data_generator

import (
	"SIlab2/game_tree"
	. "SIlab2/gameplay"
	. "SIlab2/strategies"
	"fmt"
	"math"
)

var wonConfig = map[[2]int]bool{
	{15, 15}: true,
	{15, 14}: true,
	{15, 13}: true,
	{15, 12}: true,
	{15, 11}: true,
	{14, 15}: true,
	{14, 14}: true,
	{14, 13}: true,
	{14, 12}: true,
	{14, 11}: true,
	{13, 15}: true,
	{13, 14}: true,
	{13, 13}: true,
	{13, 12}: true,
	{12, 15}: true,
	{12, 14}: true,
	{12, 13}: true,
	{11, 15}: true,
	{11, 14}: true,
}

var board = [][]int{
	{1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 2},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 2, 2},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 2, 2, 2},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 2, 2, 2, 2},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 2, 2, 2, 2},
}

var pointBoard = [][]int{
	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
	{1, 1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
	{2, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17},
	{3, 3, 4, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18},
	{4, 4, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19},
	{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
	{6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21},
	{7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22},
	{8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23},
	{9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
	{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25},
	{11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 26, 26},
	{12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 26, 27, 27},
	{13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 26, 27, 28, 28},
	{14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 26, 27, 28, 29, 29},
	{15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
}

func copyBoard(board [][]int) [][]int {
	newBoard := make([][]int, len(board))
	for i := range board {
		newBoard[i] = make([]int, len(board[i]))
		copy(newBoard[i], board[i])
	}
	return newBoard
}

func returnPositionFitness(board [][]int) float64 {
	fitness := 0.0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 1 {
				fitness += float64(pointBoard[i][j])
			} else if board[i][j] == 2 {
				fitness -= float64(30-pointBoard[i][j]) * 0.75
			}
		}
	}

	return fitness
}

func makeMoves(movesNumber int, playerState, opponentState, crowdedState float64, board *[][]int) (float64, *[][]int) {
	play1 := NewGamePlay(copyBoard(*board),
		1,
		3,
		&MainStrategy{
			playerState,
			opponentState,
			crowdedState,
			0.25},
		wonConfig,
		true)

	play2 := NewGamePlay(copyBoard(*board),
		2,
		3,
		&DistanceStrategy{},
		wonConfig,
		true)

	var move1 *game_tree.Node
	var move2 *game_tree.Node
	for i := 0; i < movesNumber; i++ {
		move1 = play1.MakeMove(0)
		if move1.AssertEquals(0, 0, 0, 0) {
			break
		}
		play2.OpponentMove(move1.StartX, move1.StartY, move1.EndX, move1.EndY)
		move2 = play2.MakeMove(0)
		if move2.AssertEquals(0, 0, 0, 0) {
			break
		}
		play1.OpponentMove(move2.StartX, move2.StartY, move2.EndX, move2.EndY)
	}

	return returnPositionFitness(*play1.GetBoard()), play1.GetBoard()
}

func Generate(step float64, maxIterations, deep int, board *[][]int) (float64, float64, float64, *[][]int) {
	playerState := 0.75
	opponentState := -0.5
	crowdedState := 0.25
	var newBoard *[][]int
	var result float64
	for i := 0; i < maxIterations; i++ {

		currentFitness, _ := makeMoves(deep, playerState, opponentState, crowdedState, board)

		nextXFitness, _ := makeMoves(deep, playerState+step, opponentState, crowdedState, board)
		previousXFitness, _ := makeMoves(deep, playerState-step, opponentState, crowdedState, board)

		if nextXFitness > currentFitness {
			playerState += step
		} else if previousXFitness > currentFitness {
			playerState -= step
		}

		nextYFitness, _ := makeMoves(deep, playerState, opponentState+step, crowdedState, board)
		previousYFitness, _ := makeMoves(deep, playerState, opponentState-step, crowdedState, board)

		if nextYFitness > currentFitness {
			opponentState += step
		} else if previousYFitness > currentFitness {
			opponentState -= step
		}

		nextZFitness, nextNewBoardZ := makeMoves(deep, playerState, opponentState, crowdedState+step, board)
		previousZFitness, prevNewBoard := makeMoves(deep, playerState, opponentState, crowdedState-step, board)

		if nextZFitness > currentFitness {
			crowdedState += step
			newBoard = nextNewBoardZ
			result = nextZFitness
		} else if previousZFitness > currentFitness {
			crowdedState -= step
			newBoard = prevNewBoard
			result = previousZFitness
		}

		if math.Abs(result-currentFitness) < 0.01 {
			break
		}
	}

	return playerState, opponentState, crowdedState, newBoard
}

func GenerateThroughMoves(step float64, maxIterations, deep, stopMovesNumber, populationSize int) *[]string {
	var playerState float64
	var opponentState float64
	var crowdedState float64
	var boards = []*[][]int{}
	var temporaryBoard *[][]int
	var lines []string
	lines = append(lines, "moveNumber,playerState,opponentState,crowdedState")
	for i := 0; i < populationSize; i++ {
		fmt.Printf("-----------------------------------%d iteracja------------------------------------------------\n", i)
		playerState, opponentState, crowdedState, temporaryBoard = Generate(step, maxIterations, deep, &board)
		boards = append(boards, temporaryBoard)

		lines = append(lines, fmt.Sprintf("%f,%f,%f,%f", i*deep, playerState, opponentState, crowdedState))
	}

	for i := 0; i < stopMovesNumber-1; i++ {
		fmt.Printf("-----------------------------------%d stop move num------------------------------------------------\n", i)
		var newTemporaryBoards []*[][]int
		for _, newBoard := range boards {
			fmt.Printf("__________________________________________%d stop move num__________________________________________\n", i)
			playerState, opponentState, crowdedState, temporaryBoard = Generate(step, maxIterations, deep, newBoard)

			newTemporaryBoards = append(newTemporaryBoards, temporaryBoard)
			lines = append(lines, fmt.Sprintf("%d,%f,%f,%f", i*deep, playerState, opponentState, crowdedState))
		}

		boards = newTemporaryBoards
	}

	return &lines
}
