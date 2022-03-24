package lemin

import "fmt"

// AddVertex adds a Vertex to the Graph
func (g *Graph) AddVertex(k string) error {

	if Contains(g.Rooms, k) {
		err := fmt.Errorf("ERROR: room %v is already exists", k)
		return err
	} else {
		g.Rooms = append(g.Rooms, &Room{Key: k})
	}
	return nil
}

// Adds an edge to the graph
func (g *Graph) AddEdge(from, to string) error {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	// error handling
	if fromVertex == nil || toVertex == nil || fromVertex == toVertex {
		err := fmt.Errorf("ERROR: invalid edge(%v --> %v)", from, to)
		return err
	} else if Contains(fromVertex.Children, to) {
		err := fmt.Errorf("ERROR: existing edge(%v --> %v)", from, to)
		return err
	} else {
		// add the edge
		fromVertex.Children = append(fromVertex.Children, toVertex)
		toVertex.Children = append(toVertex.Children, fromVertex)
	}
	return nil
}

// Returns a pointer to the Vertex with a key
func (g *Graph) getVertex(k string) *Room {
	for i, v := range g.Rooms {
		if v.Key == k {
			return g.Rooms[i]
		}
	}
	return nil
}

// Checks the adding vertex for existency
func Contains(s []*Room, key string) bool {
	for _, v := range s {
		if v.Key == key {
			return true
		}
	}
	return false
}