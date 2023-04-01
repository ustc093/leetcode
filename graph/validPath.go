package graph

//有一个具有 n 个顶点的 双向 图，其中每个顶点标记从 0 到 n - 1（包含 0 和 n - 1）。图中的边用一个二维整数数组 edges 表示，其中
//edges[i] = [ui, vi] 表示顶点 ui 和顶点 vi 之间的双向边。 每个顶点对由 最多一条 边连接，并且没有顶点存在与自身相连的边。
//
// 请你确定是否存在从顶点 source 开始，到顶点 destination 结束的 有效路径 。
//
// 给你数组 edges 和整数 n、source 和 destination，如果从 source 到 destination 存在 有效路径 ，则返回
//true，否则返回 false 。
//
//
//
// 示例 1：
//
//
//输入：n = 3, edges = [[0,1],[1,2],[2,0]], source = 0, destination = 2
//输出：true
//解释：存在由顶点 0 到顶点 2 的路径:
//- 0 → 1 → 2
//- 0 → 2
//
//
// 示例 2：
//
//
//输入：n = 6, edges = [[0,1],[0,2],[3,5],[5,4],[4,3]], source = 0, destination = 5
//
//输出：false
//解释：不存在由顶点 0 到顶点 5 的路径.
//
//
//
//
// 提示：
//
//
// 1 <= n <= 2 * 10⁵
// 0 <= edges.length <= 2 * 10⁵
// edges[i].length == 2
// 0 <= ui, vi <= n - 1
// ui != vi
// 0 <= source, destination <= n - 1
// 不存在重复边
// 不存在指向顶点自身的边
var visited = make(map[*Node]bool)

func validPath(n int, edges [][]int, source int, destination int) bool {
	graph := genGraph(n, edges)
	sourceNode := graph.Node[source]
	destNode := graph.Node[destination]
	return helper(sourceNode, destNode)
}

func genGraph(n int, edges [][]int) *Graph {
	node := make(map[int]*Node)
	for i := 0; i < n; i++ {
		node[i] = &Node{
			Value:     i,
			NextNodes: make([]*Node, 0),
		}
	}

	for _, edge := range edges {
		sourceNode := node[edge[0]]
		destNode := node[edge[1]]

		sourceNode.NextNodes = append(sourceNode.NextNodes, destNode)
		sourceNode.Out++
		destNode.In++

		// 无向图
		destNode.NextNodes = append(destNode.NextNodes, sourceNode)
		destNode.Out++
		sourceNode.In++
	}

	// 没用到边信息
	return &Graph{
		Node: node,
	}
}

func helper(sourceNode, destNode *Node) bool {
	if sourceNode == nil || visited[sourceNode] {
		return false
	}

	if sourceNode.Value == destNode.Value {
		return true
	}
	visited[sourceNode] = true
	nextNodes := sourceNode.NextNodes
	for _, nextNode := range nextNodes {
		if helper(nextNode, destNode) {
			return true
		}
	}

	return false
}
