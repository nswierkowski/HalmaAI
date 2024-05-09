package strategies

import "math"

var board = [][]int{
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

func manhattanDistance(x1, y1, x2, y2 int) float64 {
	return math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2))
}

func returnDestination(length, player int) (int, int, int, int) {
	coordinateX := 0
	coordinateY := 0
	enemiesCoordinateX := length - 1
	enemiesCoordinateY := length - 1
	if player != 2 {
		coordinateX = enemiesCoordinateX
		coordinateY = enemiesCoordinateY
		enemiesCoordinateX = 0
		enemiesCoordinateY = 0
	}
	return coordinateX, coordinateY, enemiesCoordinateX, enemiesCoordinateY
}

func countManhattanDistanceSum(playerPawns map[[2]int]bool, length, currentPlayer int) float64 {
	playerX, playerY, _, _ := returnDestination(length, currentPlayer)
	playerState := 0.0
	for pawn := range playerPawns {
		playerState += 30 - manhattanDistance(pawn[0], pawn[1], playerX, playerY)
	}
	return playerState
}

func countPlayerResults(pawns map[[2]int]bool, currentPlayer int) float64 {
	playerState := 0.0
	for pawn := range pawns {
		if currentPlayer != 1 {
			playerState += float64(board[pawn[0]][pawn[1]])
		} else {
			playerState += 30.0 - float64(board[pawn[0]][pawn[1]])
		}
	}
	return playerState
}

func countOpponentsResults(pawns map[[2]int]bool, currentPlayer int) float64 {
	playerState := 0.0
	for pawn := range pawns {
		if currentPlayer == 1 {
			playerState += float64(board[pawn[0]][pawn[1]])
		} else {
			playerState += 30.0 - float64(board[pawn[0]][pawn[1]])
		}
	}
	return playerState
}
