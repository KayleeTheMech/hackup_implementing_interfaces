package main

import (
	"gonum.org/v1/gonum/graph"
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

}

func (g Graph) From(id int64) graph.Nodes {

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
