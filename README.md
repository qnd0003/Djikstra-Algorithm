# Djikstra Algorithm

Shortest path between two nodes using Djikstra Algorithm.

### Data

This will take in an undirected graph file with three fields. The first two fields are the nodes and the last field is the vertices between the two nodes. See [example graph](###Example file) below.

## Example file

_graph1.txt_

| Node 1 | Node 2 | Vertices |
| :----- | ------ | -------- |
| A      | C      | 10       |
| C      | D      | 19       |
| A      | B      | 9        |

_No need to repeat the vertices_

## Usage

###### Compile and excute

```go
go run main.go ./graph1.txt
```

###### User input

```
source: a
target: c
```

## Sources

[Djikstra Shortest Path](https://www.geeksforgeeks.org/dijkstras-shortest-path-algorithm-greedy-algo-7/)















