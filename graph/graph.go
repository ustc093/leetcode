package graph

type Node struct {
	Value     int
	In        int
	Out       int
	NextNodes []*Node
	OutEdges  []*Edge
}

type Edge struct {
	Weight   int
	FromNode int
	ToNode   int
}

type Graph struct {
	Node map[int]*Node
	Edges []*Edge
}
