package graph

import "testing"

func TestHelper(t *testing.T) {
	Node0 := &Node{
		Value: 0,
	}

	Node1 := &Node{
		Value: 1,
	}

	Node2 := &Node{
		Value: 2,
	}

	Node3 := &Node{
		Value: 3,
	}

	Node0.NextNodes = []*Node{
		Node1, Node2,
	}

	Node1.NextNodes = []*Node{
		Node2,
	}

	Node2.NextNodes = []*Node{
		Node3,
	}

	res := helper(Node2, Node1)
	println(res)
}

func TestValidPath(t *testing.T) {
	re := validPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 3, 5)
	println(re)
}
