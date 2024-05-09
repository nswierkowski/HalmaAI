package main

import (
	. "SIlab2/data_generator"
	. "SIlab2/endpoints"
	. "SIlab2/game_tree"
	. "SIlab2/gameplay"
	. "SIlab2/strategies"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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

var board2 = [][]int{
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

func writeLinesToFile(lines []string, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func countTime(gameplay_1, gameplay_2 *GamePlay) {
	var move1 *Node
	var move2 *Node
	var lines []string
	var time1 time.Time
	var time2 time.Time
	var time1_end time.Duration
	var time2_end time.Duration
	lines = append(lines, "minmax,alfabeta")
	for gameplay_1.IsGameEnded() || gameplay_2.IsGameEnded() {
		time1 = time.Now()
		move1 = gameplay_1.MakeMove(1)
		time1_end = time.Since(time1)
		if move1.AssertEquals(0, 0, 0, 0) {
			break
		}
		gameplay_2.OpponentMove(move1.StartX, move1.StartY, move1.EndX, move1.EndY)
		gameplay_2.PrintBoard()
		time2 = time.Now()
		move2 = gameplay_2.MakeMove(2)
		time2_end = time.Since(time2)
		if move2.AssertEquals(0, 0, 0, 0) {
			break
		}
		gameplay_1.OpponentMove(move2.StartX, move2.StartY, move2.EndX, move2.EndY)
		gameplay_1.PrintBoard()

		lines = append(lines, fmt.Sprintf("%d,%d", time1_end.Nanoseconds(), time2_end.Nanoseconds()))
	}
	gameplay_1.PrintBoard()

	err := writeLinesToFile(lines, "./results/time_compare.csv")
	if err != nil {
		return
	}
}

func getUserMove(board [][]int) (int, int, int, int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter four integers separated by spaces: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, 0, 0, 0, err
	}

	input = strings.TrimSpace(input)
	values := strings.Split(input, " ")
	if len(values) != 4 {
		return 0, 0, 0, 0, fmt.Errorf("expected four integers separated by spaces")
	}

	var a, b, c, d int
	a, err = strconv.Atoi(values[0])
	if err != nil {
		return 0, 0, 0, 0, err
	}
	b, err = strconv.Atoi(values[1])
	if err != nil {
		return 0, 0, 0, 0, err
	}
	c, err = strconv.Atoi(values[2])
	if err != nil {
		return 0, 0, 0, 0, err
	}
	d, err = strconv.Atoi(values[3])
	if err != nil {
		return 0, 0, 0, 0, err
	}

	if board[a][b] == 0 || board[c][d] != 0 {
		return 0, 0, 0, 0, fmt.Errorf("illegal move")
	}

	return a, b, c, d, nil
}

func playWithRealPlayer() {
	var gameplay_1 = NewGamePlay(board,
		2,
		3,
		&DistanceStrategy{},
		wonConfig,
		true)
	gameplay_1.PrintBoard()
	startX, startY, endX, endY, err := getUserMove(board)
	for gameplay_1.IsGameEnded() {
		for err != nil {
			startX, startY, endX, endY, err = getUserMove(board)
		}
		gameplay_1.OpponentMove(startX, startY, endX, endY)
		gameplay_1.PrintBoard()
		gameplay_1.MakeMove(0)
		gameplay_1.PrintBoard()
		startX, startY, endX, endY, err = getUserMove(board)
	}
	gameplay_1.PrintBoard()

}

func hillClimbingCount() {
	err := writeLinesToFile(*GenerateThroughMoves(0.1, 10, 10, 10, 3), "./results/best_coefficient.csv")
	if err != nil {
		return
	}
}

func defaultGame() {
	var gameplay_1 = NewGamePlay(board,
		1,
		3,
		&AdaptiveStrategy{CoefficientManhattanSum: 0.25},
		wonConfig,
		true)
	var gameplay_2 = NewGamePlay(board2,
		2, 3,
		&MainStrategy{CoefficientPlayerState: 0.75, CoefficientOpponentState: -0.5, CoefficientPlayerCrowded: 0.25, CoefficientManhattanSum: 0.25},
		wonConfig,
		true)

	var move1 *Node
	var move2 *Node
	var time1 time.Time
	var time2 time.Time
	var time1_end time.Duration
	var time2_end time.Duration
	for gameplay_1.IsGameEnded() || gameplay_2.IsGameEnded() {
		time1 = time.Now()
		move1 = gameplay_1.MakeMove(1)
		time1_end = time.Since(time1)
		gameplay_1.PrintBoard()
		fmt.Fprintln(os.Stderr, "Move execution time: ", time1_end.Milliseconds())
		fmt.Fprintln(os.Stderr, "Visited nodes: ", gameplay_2.NodesVisited)
		if move1.AssertEquals(0, 0, 0, 0) {
			break
		}
		gameplay_2.OpponentMove(move1.StartX, move1.StartY, move1.EndX, move1.EndY)
		time2 = time.Now()
		move2 = gameplay_2.MakeMove(2)
		time2_end = time.Since(time2)
		gameplay_2.PrintBoard()
		fmt.Fprintln(os.Stderr, "Move execution time: ", time2_end.Milliseconds())
		fmt.Fprintln(os.Stderr, "Visited nodes: ", gameplay_2.NodesVisited)
		if move2.AssertEquals(0, 0, 0, 0) {
			break
		}
		gameplay_1.OpponentMove(move2.StartX, move2.StartY, move2.EndX, move2.EndY)

	}
}

func runServer() {
	listener := Listener{
		NewGamePlay(board, 2, 3, &DistanceStrategy{}, wonConfig, true)}
	listener.Run()
}

func runClient() {
	var gameplay_1 = NewGamePlay(board,
		1,
		3,
		&DistanceStrategy{},
		wonConfig,
		true)

	var move1 *Node
	var move2 Node
	var time1 time.Time
	var time1_end time.Duration
	for gameplay_1.IsGameEnded() {
		time1 = time.Now()
		move1 = gameplay_1.MakeMove(1)
		time1_end = time.Since(time1)
		gameplay_1.PrintBoard()
		fmt.Fprintln(os.Stderr, "Move execution time: ", time1_end.Milliseconds())
		fmt.Fprintln(os.Stderr, "Visited nodes: ", gameplay_1.NodesVisited)
		if move1.AssertEquals(0, 0, 0, 0) {
			break
		}

		move2 = SentMove(*move1)
		if move2.AssertEquals(0, 0, 0, 0) {
			break
		}
		gameplay_1.OpponentMove(move2.StartX, move2.StartY, move2.EndX, move2.EndY)

	}
	gameplay_1.PrintBoard()
}

func main() {
	defaultGame()
}
