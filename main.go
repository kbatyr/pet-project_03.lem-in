package main

import (
	"fmt"
	lemin "lem-in/src"
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("ERROR: invalid number of arguments in terminal")
		return
	}

	// Reading file
	arrFile, err := lemin.ReadFile(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get input data
	input := &lemin.Input{}
	if err := input.GetData(arrFile); err != nil {
		fmt.Println(err)
		return
	}

	// Adding rooms
	graph := &lemin.Graph{}
	for _, v := range input.Rooms {
		if err := graph.AddVertex(v); err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	// Adding tunnels (links) between rooms
	for _, v := range input.Links {
		if err := graph.AddEdge(v[0], v[1]); err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	// Finding all possible paths from start to end room
	allPaths := &lemin.Allpaths{}
	allPaths.Paths = graph.BFS(input.StartR, input.EndR)

	if len(allPaths.Paths) == 0 {
		fmt.Println("There is no path between Start and End rooms")
		return
	}

	// Finding all combinations of paths that don't intersects with each other
	allPaths.Combinations()

	// Calculating the time of moving of ants from start to end room for each combination
	ticks := allPaths.Ticks(input.Ants)

	// Choosing the best combination based on num of ants and ticks
	allPaths.OptimalPath(ticks, graph)

	// Calculating of the num of ants for each path of combination
	allPaths.AntsAllocation(input.Ants)

	// Printing the migration of ants with only start and end rooms (only 1 step migration)
	if allPaths.OptPath[0].P[0] == graph.Start && allPaths.OptPath[0].P[1] == graph.End {

		allPaths.PrintOneStep(input.Ants)
	} else {

		// Printing the migration of ants along the paths
		allPaths.Output(input.Ants)
	}
}
