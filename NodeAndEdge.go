package main

import "gonum.org/v1/gonum/graph"

type Node struct {
	NodeName string
	id       int64
}

func NewNode(name string, identity int64) Node {
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
	from     Node
	to       Node
}

func NewEdge(name string, edgeWeight float64, fromNode Node, toNode Node) Edge {
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

func (e Edge) To()  graph.Node  {
	return e.to
}

func (e Edge) Weight() float64 {
	return e.weight
}
