package main

import (
	"fmt"
	"gonum.org/v1/gonum/graph/path"
)

func main() {
	fmt.Println("My own pathfinder")
	g := NewGraph()
	node1 := NewNode("Erste", 1)
	g.AddNode(node1)
	node2 := NewNode("Zweite", 2)
	g.AddNode(node2)
	node3 := NewNode("Dritte", 3)
	g.AddNode(node3)
	node4 := NewNode("Vierte", 4)
	g.AddNode(node4)
	node5 := NewNode("Fünfte", 5)
	g.AddNode(node5)
	node6 := NewNode("Sechste", 6)
	g.AddNode(node6)

	g.SetEdge(NewEdge("erste kante", 0.5, node1, node2))
	g.SetEdge(NewEdge("zweite kante", 4.5, node1, node3))
	g.SetEdge(NewEdge("dritte kante", 0.5, node3, node4))
	g.SetEdge(NewEdge("vierte kante", 0.5, node4, node5))
	g.SetEdge(NewEdge("fünfte kante", 0.5, node4, node6))

	test, _ := path.AStar(node1, node6, g, nil)

	path, distance := test.To(node6.id)
	for idx, node := range path {
		fmt.Println(idx+1, "Step has ID", node.ID())
	}
	fmt.Println("Total distance to target node was", distance)
}
