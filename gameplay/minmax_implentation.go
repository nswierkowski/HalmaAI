package gameplay

import (
	. "SIlab2/game_tree"
	"math"
)

func (thisGame *GamePlay) createChild(move [2]int, node *Node, x, y int, board [][]int, pawns, enemiesPawns map[[2]int]bool, player int) *Node {
	moveX, moveY := move[0], move[1]
	child := &Node{
		StartX:    x,
		StartY:    y,
		EndX:      moveX,
		EndY:      moveY,
		Parent:    node,
		Childless: thisGame.assertGameEndedByBoard(board),
		Fitness:   thisGame.determinesGameState.Evaluate(pawns, enemiesPawns, len(board), player, thisGame.player),
	}
	node.AddChild(child)
	return child
}

func (thisGame *GamePlay) createNewChildrenMin(node *Node, board [][]int, pawns, enemiesPawns map[[2]int]bool, depth,
	player int, alfa, beta float64) (float64, float64, *Node, float64) {
	var value = float64(math.MaxInt)
	var lastChild *Node
	var isDone = false
	for pawn := range enemiesPawns {
		x, y := pawn[0], pawn[1]
		for move := range thisGame.generatePotentialMoves(x, y) {
			insertMoveByCoordinates(x, y, move[0], move[1], board, enemiesPawns)
			child := thisGame.createChild(move, node, x, y, board, enemiesPawns, pawns, 3-player)
			thisGame.NodesVisited += 1
			_, newValue := thisGame.minmax(child, board, enemiesPawns, pawns, depth+1, 3-player, alfa, beta)
			if newValue < value {
				value = newValue
				lastChild = child
			}
			moveBack(child, board, enemiesPawns)
			if value < alfa && thisGame.alfaBeta {
				isDone = true
				break
			}
			beta = min(beta, value)
			//lastChild = child
		}
		if isDone {
			isDone = false
			break
		}
	}

	node.PreferredChild = lastChild
	return alfa, beta, lastChild, value
}

func (thisGame *GamePlay) createNewChildrenMax(node *Node, board [][]int, pawns, enemiesPawns map[[2]int]bool, depth,
	player int, alfa, beta float64) (float64, float64, *Node, float64) {
	//fmt.Printf("usingExistingChildrenMin DEPTH: %d\n", depth)
	value := -float64(math.MaxInt)
	var lastChild *Node
	var isDone = false
	for pawn := range enemiesPawns {
		x, y := pawn[0], pawn[1]
		for move := range thisGame.generatePotentialMoves(x, y) {
			insertMoveByCoordinates(x, y, move[0], move[1], board, enemiesPawns)
			child := thisGame.createChild(move, node, x, y, board, enemiesPawns, pawns, 3-player)
			thisGame.NodesVisited += 1
			//insertMove(child, board, enemiesPawns)
			_, newValue := thisGame.minmax(child, board, enemiesPawns, pawns, depth+1, 3-player, alfa, beta)
			//fmt.Printf("createNewChildrenMax newValue: %f, value: %f, DEPTH: %d\n", newValue, value, depth)
			if newValue > value {
				//fmt.Printf("NEW VALUE IS BIGGER")
				value = newValue
				lastChild = child
			}
			moveBack(child, board, enemiesPawns)
			if value > beta && thisGame.alfaBeta {
				isDone = true
				break
			}
			alfa = max(alfa, value)
		}
		if isDone {
			isDone = false
			break
		}
	}

	node.PreferredChild = lastChild
	return alfa, beta, lastChild, value
}

func (thisGame *GamePlay) usingExistingChildrenMin(node *Node, board [][]int, pawns, enemiesPawns map[[2]int]bool, depth,
	player int, alfa, beta float64) (float64, float64, *Node, float64) {
	//fmt.Printf("usingExistingChildrenMin DEPTH: %d\n", depth)
	var value = float64(math.MaxInt)
	var lastChild *Node
	for child := range node.Children {
		insertMove(child, board, enemiesPawns)
		_, newValue := thisGame.minmax(child, board, enemiesPawns, pawns, depth+1, 3-player, alfa, beta)
		//fmt.Printf("usingExistingChildrenMin newValue: %f, value: %f, DEPTH: %d\n", newValue, value, depth)
		if newValue < value {
			//fmt.Printf("VALUE IS BIGGER")
			value = newValue
			lastChild = child
		}
		moveBack(child, board, enemiesPawns)
		if value < alfa && thisGame.alfaBeta {
			break
		}
		beta = min(beta, value)
		//lastChild = child
	}

	node.PreferredChild = lastChild
	return alfa, beta, lastChild, value
}

func (thisGame *GamePlay) usingExistingChildrenMax(node *Node, board [][]int, pawns, enemiesPawns map[[2]int]bool, depth,
	player int, alfa, beta float64) (float64, float64, *Node, float64) {
	//fmt.Printf("usingExistingChildrenMax DEPTH: %d\n", depth)
	var value = -float64(math.MaxInt)
	var lastChild *Node
	for child := range node.Children {
		insertMove(child, board, enemiesPawns)
		_, newValue := thisGame.minmax(child, board, enemiesPawns, pawns, depth+1, 3-player, alfa, beta)
		//fmt.Printf("usingExistingChildrenMax newValue: %f, value: %f, DEPTH: %d\n", newValue, value, depth)
		if newValue > value {
			//fmt.Printf("NEW VALUE IS BIGGER")
			value = newValue
			lastChild = child
		}
		moveBack(child, board, enemiesPawns)
		if value > beta && thisGame.alfaBeta {
			break
		}
		alfa = max(alfa, value)
	}

	node.PreferredChild = lastChild
	return alfa, beta, lastChild, value
}

func (thisGame *GamePlay) minmax(node *Node, board [][]int, pawns, enemiesPawns map[[2]int]bool, depth,
	player int, alfa, beta float64) (*Node, float64) {

	if depth >= thisGame.depth || node.Childless {
		if node.ChildFitness != 0 {
			return node, node.ChildFitness
		}
		return node, node.Fitness
	}

	var lastChild *Node
	var value float64
	if node.FitnessCounterMin {
		if len(node.Children) == 0 {
			//fmt.Printf("LEN=0 MIN DEPTH: %d\n", depth)
			alfa, beta, lastChild, value = thisGame.createNewChildrenMin(node, board, pawns, enemiesPawns, depth, player, alfa, beta)
		} else {
			//fmt.Printf("LEN!=0 MIN DEPTH: %d\n", depth)
			alfa, beta, lastChild, value = thisGame.usingExistingChildrenMin(node, board, pawns, enemiesPawns, depth, player, alfa, beta)
		}
	} else {
		if len(node.Children) == 0 {
			//fmt.Printf("LEN=0 MAX DEPTH: %d\n", depth)
			alfa, beta, lastChild, value = thisGame.createNewChildrenMax(node, board, pawns, enemiesPawns, depth, player, alfa, beta)
		} else {
			//fmt.Printf("LEN!=0 MAX DEPTH: %d\n", depth)s
			alfa, beta, lastChild, value = thisGame.usingExistingChildrenMax(node, board, pawns, enemiesPawns, depth, player, alfa, beta)
		}
	}

	node.PreferredChild = lastChild
	node.ChildFitness = value
	//fmt.Printf("VALUE: %f DEPTH: %d\n", node.ChildFitness, depth)
	return lastChild, value
}
