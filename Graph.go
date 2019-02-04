package main

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/iterator"
	"math"
)

type Graph struct {
	myNodes []Node
	myEdges []Edge
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

func (g Graph) Edge(uid, vid int64) graph.WeightedEdge {
	for _, edge := range g.myEdges {
		if edge.From().ID() == uid && edge.To().ID() == vid {
			return edge
		}
	}
	return nil
}

func (g Graph) WeightedEdge(uid, vid int64) graph.WeightedEdge {
	return g.Edge(uid, vid)
}

func (g Graph) Weight(xid, yid int64) (w float64, ok bool) {
	edge := g.WeightedEdge(xid, yid)
	if edge != nil {
		return edge.Weight(), true
	}
	return math.Inf(1), false
}
