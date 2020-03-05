package main

import (
	"fmt"
	"./functions"
	"os"
	"strings"
)

func main() {
	// command line arguments
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Please enter file name")
	}
	
	// creating the graph
	vertices := f.ReadFile(args[0])

	// print graph
	f.PrintGraph(vertices)

	// Djikstra
	for {
		// second try
		var source string
		var target string
		fmt.Print("Source: ")
		fmt.Scan(&source)
		fmt.Print("Target: ")
		fmt.Scan(&target)

		// djikstra algorithm
		f.Djikstra(vertices, strings.ToUpper(source), strings.ToUpper(target))
		// reset
		vertices = f.ResetGraph(vertices)
	}		
}