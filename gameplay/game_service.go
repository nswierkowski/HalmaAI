package gameplay

import (
	. "SIlab2/game_tree"
	. "SIlab2/strategies"
	"fmt"
	"math"
)

var winConfig = map[int]map[[2]int]bool{
	1: {
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
	},
	2: {
		{0, 0}: true,
		{0, 1}: true,
		{0, 2}: true,
		{0, 3}: true,
		{0, 4}: true,
		{1, 0}: true,
		{1, 1}: true,
		{1, 2}: true,
		{1, 3}: true,
		{1, 4}: true,
		{2, 0}: true,
		{2, 1}: true,
		{2, 2}: true,
		{2, 3}: true,
		{3, 0}: true,
		{3, 1}: true,
		{3, 2}: true,
		{4, 0}: true,
		{4, 1}: true,
	},
}

func assertCoordinatesAreValid(x, y int, board [][]int) bool {
	if x < 0 || x >= len(board) {
		return false
	}
	if y < 0 || y >= len(board[x]) {
		return false
	}
	return true
}

func getPawnsSet(board [][]int, player int) (map[[2]int]bool, map[[2]int]bool) {
	playerPawns := make(map[[2]int]bool)
	enemiesPawns := make(map[[2]int]bool)
	for x := range board {
		for y := range board[x] {
			if board[x][y] == player {
				playerPawns[[2]int{x, y}] = true
			} else if board[x][y] != 0 {
				enemiesPawns[[2]int{x, y}] = true
			}
		}
	}
	return playerPawns, enemiesPawns
}

func insertMove(node *Node, board [][]int, pawns map[[2]int]bool) {
	if node == nil {
		fmt.Println("NODE IS NIL WRRR\n")

		fmt.Println("###########################################################################")
		fmt.Println("Board:")
		for i := 0; i < 16; i++ {
			for j := 0; j < 16; j++ {
				fmt.Printf("%d ", board[i][j])
			}
			fmt.Println()
		}

		fmt.Println("###########################################################################")
		fmt.Println()
	}
	board[node.EndX][node.EndY] = board[node.StartX][node.StartY]
	board[node.StartX][node.StartY] = 0
	delete(pawns, [2]int{node.StartX, node.StartY})
	pawns[[2]int{node.EndX, node.EndY}] = true
}

func insertMoveByCoordinates(originX, originY, finalX, finalY int, board [][]int, pawns map[[2]int]bool) {
	board[finalX][finalY] = board[originX][originY]
	board[originX][originY] = 0
	delete(pawns, [2]int{originX, originY})
	pawns[[2]int{finalX, finalY}] = true
}

func moveBack(node *Node, board [][]int, pawns map[[2]int]bool) {
	pawns[[2]int{node.StartX, node.StartY}] = true
	delete(pawns, [2]int{node.EndX, node.EndY})
	board[node.StartX][node.StartY] = board[node.EndX][node.EndY]
	board[node.EndX][node.EndY] = 0
}

func moveBackByCoordinates(originX, originY, finalX, finalY int, board [][]int, pawns map[[2]int]bool) {
	pawns[[2]int{originX, originY}] = true
	delete(pawns, [2]int{finalX, finalY})
	board[originX][originY] = board[finalX][finalY]
	board[finalX][finalY] = 0
}

type GamePlay struct {
	board               [][]int
	player              int
	depth               int
	root                *Node
	determinesGameState Strategy
	wonConfiguration    map[[2]int]bool
	gameOn              bool
	alfaBeta            bool
	NodesVisited        int
}

func (thisGame *GamePlay) PrintBoard() {
	for _, row := range thisGame.board {
		for _, col := range row {
			fmt.Printf(" %d", col)
		}
		fmt.Println()
	}
}

func (thisGame *GamePlay) answerWhoseTurn() int {
	moveOfPlayer := 0
	for _, row := range thisGame.board {
		for _, col := range row {
			if col == 1 {
				moveOfPlayer++
			} else if col == 2 {
				moveOfPlayer--
			}
		}
	}
	if moveOfPlayer >= 0 {
		return 1
	}
	return 2
}

func (thisGame *GamePlay) assertJumpPossibility(prevX, prevY, x, y int) [][2]int {
	jumps := make([][2]int, 0)
	for _, newX := range []int{x - 1, x, x + 1} {
		for _, newY := range []int{y - 1, y, y + 1} {
			if (newX == prevX && newY == prevY) || !assertCoordinatesAreValid(newX, newY, thisGame.board) {
				continue
			}
			if thisGame.board[newX][newY] != 0 {
				jumps = append(jumps, [2]int{newX, newY})
			}
		}
	}
	return jumps
}

func (thisGame *GamePlay) generateJump(originX, originY, directionX, directionY int, foundMoves map[[2]int]bool) {
	newX := originX + ((directionX - originX) * 2)
	newY := originY + ((directionY - originY) * 2)
	if _, ok := foundMoves[[2]int{newX, newY}]; ok {
		return
	} else if assertCoordinatesAreValid(newX, newY, thisGame.board) && thisGame.board[newX][newY] == 0 {
		foundMoves[[2]int{newX, newY}] = true
		newJumps := thisGame.assertJumpPossibility(directionX, directionY, newX, newY)
		for _, jump := range newJumps {
			thisGame.generateJump(newX, newY, jump[0], jump[1], foundMoves)
		}
	}
}

func (thisGame *GamePlay) generatePotentialMoves(x, y int) map[[2]int]bool {
	yMove := []int{y - 1, y, y + 1}
	xMove := []int{x - 1, x, x + 1}

	moves := make(map[[2]int]bool)
	for _, newX := range xMove {
		for _, newY := range yMove {
			if !assertCoordinatesAreValid(newX, newY, thisGame.board) {
				continue
			}
			if thisGame.board[newX][newY] == 0 {
				moves[[2]int{newX, newY}] = true
			} else {
				thisGame.generateJump(x, y, newX, newY, moves)
			}
		}
	}
	return moves
}

func (thisGame *GamePlay) showMoves(x, y int) {
	fmt.Println(thisGame.generatePotentialMoves(x, y))
}

func (thisGame *GamePlay) assertGameEnded(pawns map[[2]int]bool, player int) bool {
	for pawn := range winConfig[player] {
		if !pawns[pawn] {
			return false
		}
	}
	return true
}

func (thisGame *GamePlay) assertGameEndedByBoard(board [][]int) bool {
	var firstWon bool = true
	for move := range winConfig[1] {
		if board[move[0]][move[1]] != 1 {
			firstWon = false
			break
		}
	}

	if firstWon {
		return true
	}

	for move := range winConfig[2] {
		if board[move[0]][move[1]] != 2 {
			return false
		}
	}
	return true
}

func (thisGame *GamePlay) addNodesChildren(node *Node, board [][]int, pawns, enemiesPawns map[[2]int]bool, depth,
	player int, executeMove bool) {
	if executeMove {
		insertMove(node, board, pawns)
	}

	if node.Childless {
		return
	}

	if len(node.Children) != 0 && depth < thisGame.depth {
		for child := range node.Children {
			if !child.Childless {
				thisGame.addNodesChildren(child, board, enemiesPawns, pawns, depth+1, 3-player, true)
				moveBack(child, board, enemiesPawns)
			}
		}
	} else if depth < thisGame.depth {
		for pawn := range enemiesPawns {
			x, y := pawn[0], pawn[1]
			for move := range thisGame.generatePotentialMoves(x, y) {
				moveX, moveY := move[0], move[1]
				insertMoveByCoordinates(x, y, moveX, moveY, board, enemiesPawns)
				node.AddChild(
					&Node{
						StartX:    x,
						StartY:    y,
						EndX:      moveX,
						EndY:      moveY,
						Parent:    node,
						Childless: thisGame.assertGameEnded(pawns, 3-player),
						Fitness:   thisGame.determinesGameState.Evaluate(pawns, enemiesPawns, len(board), player, thisGame.player),
					})
				moveBackByCoordinates(x, y, moveX, moveY, board, enemiesPawns)
			}
		}

		for child := range node.Children {
			if !child.Childless {
				thisGame.addNodesChildren(child, board, enemiesPawns, pawns, depth+1, 3-player, true)
				moveBack(child, board, enemiesPawns)
			}
		}
	}

	if depth < thisGame.depth {
		var preferredChild *Node
		var fitness float64
		for child := range node.Children {
			if preferredChild == nil {
				preferredChild = child
				fitness = child.Fitness
			}

			childFitness := child.Fitness
			if child.ChildFitness != 0 {
				childFitness = child.ChildFitness
			}

			if node.CompareFitness(childFitness, fitness) {
				fitness = childFitness
				preferredChild = child
			}
		}
		node.PreferredChild = preferredChild
		node.ChildFitness = fitness
	}
	//node.UpdateFitness(fitness, preferredChild)
}

func (thisGame *GamePlay) MakeMove(player int) *Node {
	if !thisGame.gameOn {
		fmt.Println("Game has ended")
		return &Node{}
	}

	if player == 0 {
		player = thisGame.player
	}

	if player == thisGame.player {
		thisGame.determinesGameState.ChangeCoefficients()
	}

	playersPawns, enemiesPawns := getPawnsSet(thisGame.board, player)
	if thisGame.assertGameEnded(playersPawns, player) {
		thisGame.gameOn = false
		fmt.Println("I'VE WON PLAYER: ", player, thisGame.determinesGameState.Evaluate(playersPawns, enemiesPawns, 16, player, thisGame.player))
		return &Node{}
	}

	lastChild, futureFitness := thisGame.minmax(thisGame.root,
		thisGame.board,
		enemiesPawns,
		playersPawns,
		0,
		player,
		-float64(math.MaxInt),
		float64(math.MaxInt))

	thisGame.root = lastChild
	insertMove(thisGame.root, thisGame.board, playersPawns)
	thisGame.root.Parent = nil
	fmt.Println(fmt.Sprintf("FITNESS: %f FUTURE FITNESS: %f Min: %t", thisGame.root.Fitness, futureFitness, thisGame.root.FitnessCounterMin))
	thisGame.gameOn = !thisGame.root.Childless
	if !thisGame.gameOn {
		fmt.Printf("I'VE WON PLAYER: %d\n", player)
		return &Node{}
	}
	fmt.Println("I AM NOT CHILDLESS")
	fmt.Println("SO MOVE AT LEAST: ", thisGame.root.StartX, thisGame.root.StartY, thisGame.root.EndX, thisGame.root.EndY)

	return thisGame.root
}

func (thisGame *GamePlay) OpponentMove(originX, originY, finalX, finalY int) {
	if !thisGame.gameOn {
		fmt.Println("Game has ended")
		return
	}
	var nextNode *Node
	for node := range thisGame.root.Children {
		if node.AssertEquals(originX, originY, finalX, finalY) {
			nextNode = node
			break
		}
	}

	if nextNode == nil {
		nextNode = &Node{StartX: originX, StartY: originY, EndX: finalX, EndY: finalY, Parent: nil}
	}

	nextNode.Parent = nil
	thisGame.root = nextNode
	thisGame.board[finalX][finalY] = thisGame.board[originX][originY]
	thisGame.board[originX][originY] = 0
}

func (thisGame *GamePlay) IsGameEnded() bool {
	return thisGame.gameOn
}

func (thisGame *GamePlay) GetBoard() *[][]int {
	return &thisGame.board
}

func NewGamePlay(board [][]int, player, depth int, determinesGameState Strategy, wonConfiguration map[[2]int]bool, alfaBeta bool) *GamePlay {
	return &GamePlay{
		board:               board,
		player:              player,
		depth:               depth,
		root:                &Node{},
		determinesGameState: determinesGameState,
		wonConfiguration:    wonConfiguration,
		gameOn:              true,
		alfaBeta:            alfaBeta,
	}
}
