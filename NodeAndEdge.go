package main

import "gonum.org/v1/gonum/graph"

type Node struct {
	NodeName string
	id       int64
}

func NewNode(name string, identity int64) graph.Node {
	return Node{
		NodeName: name,
		id:       identity,
	}
}

func (n Node) ID() int64 {
	return n.id
}

type Edge struct {
	EdgeName string
	weight   float64
	from     graph.Node
	to       graph.Node
}

func NewEdge(name string, edgeWeight float64, fromNode graph.Node, toNode graph.Node) graph.WeightedEdge {
	return Edge{
		EdgeName: name,
		weight:   edgeWeight,
		from:     fromNode,
		to:       toNode,
	}
}

func (e Edge) From() graph.Node {
	return e.from
}

func (e Edge) To() graph.Node {
	return e.to
}

func (e Edge) Weight() float64 {
	return e.weight
}
