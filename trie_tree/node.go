package trie_tree

type Node struct {
	Pass     int
	End      int
	Children []*Node
}

type Trie struct {
	RootNode *Node
}

func NewNode() *Node {
	return &Node{
		Pass:     0,
		End:      0,
		Children: make([]*Node, 26),
	}
}


func NewTrie() *Trie {
	return &Trie{RootNode: NewNode()}
}