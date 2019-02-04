package main

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/iterator"
	"math"
	"math/rand"
	"strconv"
)

type Graph struct {
	myNodes []Node
	myEdges []Edge
}

func NewGraph() Graph {
	g := Graph{
		myNodes: make([]Node, 0),
		myEdges: make([]Edge, 0),
	}
	return g
}

func (g Graph) Node(id int64) graph.Node {
	return NewNode("empty", 0)
}

func (g Graph) Nodes() graph.Nodes {
	returnNodes := make([]graph.Node, 0)
	for _, node := range g.myNodes {
		returnNodes = append(returnNodes, node)
	}

	return iterator.NewOrderedNodes(returnNodes)
}

func (g Graph) From(id int64) graph.Nodes {
	returnArray := make([]graph.Node, 0)
	for _, edge := range g.myEdges {
		if edge.from.ID() == id {
			returnArray = append(returnArray, edge.To())
		}
	}
	return iterator.NewOrderedNodes(returnArray)
}

func (g Graph) HasEdgeBetween(xid, yid int64) bool {
	for _, edge := range g.myEdges {
		if edge.From().ID() == xid && edge.To().ID() == yid {
			return true
		}
	}
	return false
}

func (g Graph) Edge(uid, vid int64) graph.Edge {
	return g.WeightedEdge(uid, vid)
}

func (g Graph) WeightedEdge(uid, vid int64) graph.WeightedEdge {
	for _, edge := range g.myEdges {
		if edge.From().ID() == uid && edge.To().ID() == vid {
			return edge
		}
	}
	return nil
}

func (g Graph) Weight(xid, yid int64) (w float64, ok bool) {
	edge := g.WeightedEdge(xid, yid)
	if edge != nil {
		return edge.Weight(), true
	}
	return math.Inf(1), false
}

func (g *Graph) AddNode(node Node) {
	g.myNodes = append(g.myNodes, node)
}

func (g *Graph) NewNode() Node {
	newId := rand.Int63()
	n := NewNode(strconv.Itoa(int(newId)), newId)
	g.myNodes = append(g.myNodes, n)
	return n
}

func (g *Graph) RemoveNode(id int64) {
	var idToDelete int
	for idx, node := range g.myNodes {
		if node.ID() == id {
			idToDelete = idx
		}
	}
	g.myNodes = append(g.myNodes[:idToDelete-1], g.myNodes[idToDelete+1:]...)
}

func (g *Graph) NewEdge(from, to Node) Edge {
	e := NewEdge(strconv.Itoa(int(from.ID()))+strconv.Itoa(int(to.ID())), 0, from, to)
	g.myEdges = append(g.myEdges, e)
	return e
}

func (g *Graph) SetEdge(e Edge) {
	g.myEdges = append(g.myEdges, e)
}
