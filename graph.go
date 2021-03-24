package main

type Graph interface {
	AddNode(node *GraphNode)
	AddEdge(n1, n2 *GraphNode)

	String() string
}

func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{}
}

type DirectedGraph struct {
	nodes []GraphNode
	edges []GraphEdge
}

type GraphNode struct {
	Id    string
	Value interface{}
}

type GraphEdge struct {
	From GraphNode
	To   []GraphNode
}

func (dg *DirectedGraph) AddNode(node GraphNode) {
	dg.nodes = append(dg.nodes, node)
}

func (dg DirectedGraph) AddEdge(from GraphNode, to []GraphNode) {
	dg.edges = append(dg.edges, GraphEdge{
		From: from,
		To:   to,
	})
}

func (dg DirectedGraph) String() (res string) {
	for _, node := range dg.nodes {
		res += node.Id + " -> "
	}

	return res[:len(res)-4]
}
