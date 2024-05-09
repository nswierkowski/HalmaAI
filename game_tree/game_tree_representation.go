package game_tree

type Node struct {
	StartX, StartY, EndX, EndY int
	Fitness                    float64
	Parent                     *Node
	Children                   NodesChildren
	FitnessCounterMin          bool
	Childless                  bool
	PreferredChild             *Node
	ChildFitness               float64
}

type NodesChildren map[*Node]bool

func (thisNode *Node) CompareFitness(newFitness, originalFitness float64) bool {
	var comparisonFitness = newFitness > originalFitness
	if thisNode.FitnessCounterMin {
		comparisonFitness = !comparisonFitness
	}
	return comparisonFitness
}

func (thisNode *Node) UpdateFitness(fitness float64, potentialPreferredChild *Node) {

	if thisNode.CompareFitness(fitness, thisNode.ChildFitness) {
		thisNode.ChildFitness = fitness
		thisNode.PreferredChild = potentialPreferredChild

		if thisNode.Parent != nil {
			thisNode.Parent.UpdateFitness(fitness, thisNode)
		}
	}
}

func (thisNode *Node) AddChild(node *Node) bool {
	if thisNode.Childless {
		return false
	}

	if thisNode.Children == nil {
		thisNode.Children = make(map[*Node]bool)
	}

	thisNode.Children[node] = true
	node.FitnessCounterMin = thisNode.FitnessCounterMin != true
	if len(thisNode.Children) == 1 {
		thisNode.ChildFitness = node.Fitness
		thisNode.PreferredChild = node
	}

	return thisNode.CompareFitness(node.Fitness, thisNode.Fitness)
}

func (thisNode *Node) AssertEquals(startX, startY, endX, endY int) bool {
	return thisNode.StartX == startX &&
		thisNode.StartY == startY &&
		thisNode.EndY == endY &&
		thisNode.EndX == endX
}
