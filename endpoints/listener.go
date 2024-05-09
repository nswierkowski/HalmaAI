package endpoints

import (
	"SIlab2/gameplay"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Listener struct {
	GamePlay *gameplay.GamePlay
}

type NodeDTO struct {
	StartX, StartY, EndX, EndY int
}

func (listener *Listener) Run() {
	http.HandleFunc("/play", listener.play)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (listener *Listener) play(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var nodeDTO NodeDTO
	err := json.NewDecoder(r.Body).Decode(&nodeDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	fmt.Println("Received Node:", nodeDTO)
	listener.GamePlay.OpponentMove(nodeDTO.StartX, nodeDTO.StartY, nodeDTO.EndX, nodeDTO.EndY)
	start_time := time.Now()
	move := listener.GamePlay.MakeMove(0)
	execution_time := time.Since(start_time)
	listener.GamePlay.PrintBoard()
	fmt.Fprintln(os.Stderr, "Move execution time: ", execution_time.Milliseconds())
	fmt.Fprintln(os.Stderr, "Visited nodes: ", listener.GamePlay.NodesVisited)

	w.Header().Set("Content-Type", "application/json")
	response := &NodeDTO{move.StartX, move.StartY, move.EndX, move.EndY}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
