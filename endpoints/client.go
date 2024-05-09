package endpoints

import (
	"SIlab2/game_tree"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SentMove(node game_tree.Node) game_tree.Node {
	nodeDTO := NodeDTO{node.StartX, node.StartY, node.EndX, node.EndY}
	nodeJSON, err := json.Marshal(nodeDTO)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	resp, err := http.Post("http://localhost:8080/play", "application/json", bytes.NewBuffer(nodeJSON))
	if err != nil {
		log.Fatalf("Error calling endpoint: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println(resp)
	var responseNode NodeDTO
	err = json.NewDecoder(resp.Body).Decode(&responseNode)
	if err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}

	return game_tree.Node{
		StartX: responseNode.StartX,
		StartY: responseNode.StartY,
		EndX:   responseNode.EndX,
		EndY:   responseNode.EndY}
}
