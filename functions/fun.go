package f

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"log"
	"fmt"
	"sort"
)

// Node struct
type Node struct {
	Label string
	DistantFrom int
	Adjacent[] Node
	ShortDis int
	Previous string
	Visit bool
}

// read from the file
func ReadFile(filename string) map[string] Node{
	// make vertices
	var vertex Node

	// slice of vertices
	var vertices map[string] Node
	vertices = make(map[string] Node)

	// reading from file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("error")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		b := strings.Split(scanner.Text(), " ")
		// convert
		convert_dis, err2 := strconv.Atoi(b[2])
		if err2 != nil {
			log.Fatal(err2)
		}
		
		// vertex and adj
		from := b[1]
		start := b[0]
		from = strings.ToUpper(from)
		start = strings.ToUpper(start)

		// check map and create new vertex
		if _, ok := vertices[start]; !ok {
			vertex.Label = start
			vertices[start] = vertex
		}
		if _, ok := vertices[from]; !ok {
			vertex.Label = from
			vertices[from] = vertex
		}

		// "from" is "distant" from "start"
		temp_node := vertices[from]
		temp_node.Adjacent = nil
		temp_node.DistantFrom = convert_dis
		start_node := vertices[start]
		start_node.Adjacent = Add(temp_node, vertices[start])
		vertices[start] = start_node

		// "start" is "distant" from "from" (bidirectional)
		temp_node = vertices[start]
		temp_node.Adjacent = nil
		temp_node.DistantFrom = convert_dis
		start_node = vertices[from]
		start_node.Adjacent = Add(temp_node, vertices[from])
		vertices[from] = start_node
	}
	// close
	file.Close()

	// return created graph
	return vertices
}

// fun add adj
func Add(adj Node, vertex Node) []Node{
	vertex.Adjacent = append(vertex.Adjacent, adj)
	return vertex.Adjacent
}

// print the
func PrintGraph(vertices map[string] Node) {
	// print the graph
	for _, i := range vertices {
		fmt.Println(i.Label)
		for _, j := range i.Adjacent {
			fmt.Printf("%s:%d ", j.Label, j.DistantFrom)
		}
		fmt.Println()
	}
	fmt.Println()
}

// Djikstra
func Djikstra(vertices map[string] Node, source string, target string) {
	var queue[] Node
	var visited[] string
	queue = append(queue, vertices[source])

	current := vertices[source]
	current.ShortDis = 0
	vertices = SetInfinite(vertices)

	// queue = append (queue, current)
	for target != current.Label {

		fmt.Print("Queue [ ")
		for _, x := range queue {
			fmt.Printf("%s%d ", x.Label, x.ShortDis)
		}
		fmt.Println("]")

		// pop
		queue = queue[1:]

		fmt.Println()
		fmt.Printf("Current Node: %s\n", current.Label)

		// checking adjacent
		for _, i := range current.Adjacent {
			if vertices[i.Label].Visit == false {
				calc := current.ShortDis + i.DistantFrom

				fmt.Printf("Neighbor [%s] %d+%d=%d, (%d)\n", i.Label, current.ShortDis, i.DistantFrom, calc, vertices[i.Label].ShortDis)

				// insert into the queue
				if empty := check (queue, i.Label); empty {
					a := vertices[i.Label]
					a.ShortDis = calc
					a.Previous = current.Label
					vertices[a.Label] = a

					a.Previous = current.Label
					queue = append(queue, a)
				}

				// update the node shortest distant
				if calc < vertices[i.Label].ShortDis {
					for j := 0; j < len(queue); j++ {
						if queue[j].Label == i.Label {
							queue[j].ShortDis = calc
							queue[j].Previous = current.Label
							vertices[queue[j].Label] = queue[j]
						}
					}
				}
			}
		}

		// visited
		visited = append (visited, current.Label)
		
		// pop and sort
		sort.Slice(queue, func(k, l int) bool { return queue[k].ShortDis < queue[l].ShortDis })

		fmt.Printf("shortest: %s\n\n", queue[0].Label)
		fmt.Print("visited ", visited, "\n")

		current.Visit = true
		vertices[current.Label] = current

		current = queue[0]
	}

	fmt.Println("Djikstra Completed")

	// print table
	for _, i := range vertices {
		fmt.Printf("%s %-2d %s\n", i.Label, i.ShortDis, i.Previous)
	}

	// print path
	Path(vertices, source, target)
}

// set infinite
func SetInfinite(vertices map[string] Node) map[string] Node {
	for _, i := range vertices {
		i.ShortDis = 10000
		vertices[i.Label] = i
	}
	return vertices
}

// reset of graph
func ResetGraph(vertices map[string] Node) map[string] Node{
	// reset graph
	for _, i := range vertices {
		i.ShortDis = 0
		i.Previous = " "
		i.Visit = false
		vertices[i.Label] = i
	}

	// fresh map
	return vertices
}

// path
func Path(vertices map[string] Node, source string, target string) {
	// final path
	var path[] string
	path = append(path, target)
	for target != source {
		target = vertices[target].Previous
		path = append (path, target)
	}
	
	fmt.Print("\nPath: [")
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Print(path[i])
		if i != 0 {
			fmt.Print(" - ")
		}
	}
	fmt.Println("]")
}

// check empty
func check(queue [] Node, current string) bool {
	count := 0

	for _, p := range queue {
		if current == p.Label {
			count++
		}
	}

	if count == 0 {
		return true
	} else {
		return false
	}
}