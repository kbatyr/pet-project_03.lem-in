package lemin

// Graph represents an adjacency list graph
type Graph struct {
	Rooms      []*Room
	Start, End *Room
}

// Room represents a graph vertex
type Room struct {
	Key      string
	Children []*Room
}

// Consists of rooms which have link from start to end room
type Path struct {
	P []*Room
}

// Manipulation with paths
type Allpaths struct {
	Paths    []*Path
	Combo    [][]*Path
	OptPath []*Path
	Tick     int
}

// Input data from file
type Input struct {
	Ants   int
	StartR string
	EndR   string
	Rooms  []string
	Coords [][]string
	Links  [][]string
}
