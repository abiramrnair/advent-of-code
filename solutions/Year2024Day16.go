package solutions

import (
	"aoc/utils"
	"container/heap"
	"fmt"
	"math"
)

type CoordWithDirection struct {
	utils.Coord
	Direction string
}

type Day16AdjacencyNode struct {
	CoordWithDirection CoordWithDirection
	Distance           int
}

func Day16_2024_Part1() {
	defer utils.CodeTimer()()
	INPUT_PATH := "./inputs/Year2024Day16.txt"
	input, err := utils.GetInputAs2DArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	adjacencyList := generateDay16AdjacencyList(input)
	var starting CoordWithDirection
	var ending CoordWithDirection
	ROWS, COLS := len(input), len(input[0])
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			if input[r][c] == "S" {
				starting.Coord = utils.Coord{Row: r, Col: c}
				starting.Direction = "right"
			} else if input[r][c] == "E" {
				ending.Coord = utils.Coord{Row: r, Col: c}
			}
		}
	}
	answer := djikstraDay16(adjacencyList, starting, ending)
	fmt.Println("Answer:", answer)
}

func Day16_2024_Part2() {
	defer utils.CodeTimer()()
	INPUT_PATH := "./inputs/Year2024Day16.txt"
	input, err := utils.GetInputAs2DArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	adjacencyList := generateDay16AdjacencyList(input)
	var starting CoordWithDirection
	var ending CoordWithDirection
	ROWS, COLS := len(input), len(input[0])
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			if input[r][c] == "S" {
				starting.Coord = utils.Coord{Row: r, Col: c}
				starting.Direction = "right"
			} else if input[r][c] == "E" {
				ending.Coord = utils.Coord{Row: r, Col: c}
				ending.Direction = "up"
			}
		}
	}
	answer := djikstraDay16Pt2(input, adjacencyList, starting, ending)
	fmt.Println("Answer:", answer)
}

func djikstraDay16(
	adjacencyList map[CoordWithDirection][]Day16AdjacencyNode,
	starting CoordWithDirection,
	ending CoordWithDirection,
) int {
	visited := make(map[CoordWithDirection]bool)
	distances := make(map[CoordWithDirection]int)
	for k := range adjacencyList {
		visited[k] = false
		distances[k] = math.MaxInt
	}
	pq := make(utils.PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &utils.PqItem{Value: starting, Priority: 0})
	for pq.Len() > 0 {
		currentNode := heap.Pop(&pq).(*utils.PqItem)
		currentCoord := currentNode.Value.(CoordWithDirection)
		currentDistance := currentNode.Priority
		visited[currentCoord] = true
		for _, edge := range adjacencyList[currentCoord] {
			val := visited[edge.CoordWithDirection]
			if val {
				continue
			}
			newDistance := currentDistance + edge.Distance
			if newDistance < distances[edge.CoordWithDirection] {
				distances[edge.CoordWithDirection] = newDistance
				heap.Push(&pq, &utils.PqItem{Value: edge.CoordWithDirection, Priority: newDistance})
			}
		}
	}
	minDistance := math.MaxInt
	for k, v := range distances {
		if k.Row == ending.Row && k.Col == ending.Col {
			if v < minDistance {
				minDistance = v
			}
		}
	}
	return minDistance
}

func djikstraDay16Pt2(
	grid [][]string,
	adjacencyList map[CoordWithDirection][]Day16AdjacencyNode,
	starting CoordWithDirection,
	ending CoordWithDirection,
) int {
	previous := make(map[CoordWithDirection][]CoordWithDirection)
	visited := make(map[CoordWithDirection]bool)
	distances := make(map[CoordWithDirection]int)
	for k := range adjacencyList {
		visited[k] = false
		distances[k] = math.MaxInt
		previous[k] = make([]CoordWithDirection, 0)
	}
	pq := make(utils.PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &utils.PqItem{Value: starting, Priority: 0})
	for pq.Len() > 0 {
		currentNode := heap.Pop(&pq).(*utils.PqItem)
		currentCoord := currentNode.Value.(CoordWithDirection)
		currentDistance := currentNode.Priority
		visited[currentCoord] = true
		for _, edge := range adjacencyList[currentCoord] {
			val := visited[edge.CoordWithDirection]
			if val {
				continue
			}
			newDistance := currentDistance + edge.Distance
			if newDistance <= distances[edge.CoordWithDirection] {
				distances[edge.CoordWithDirection] = newDistance
				shouldAppend := true
				for _, c := range previous[edge.CoordWithDirection] {
					if c.Coord == currentCoord.Coord && c.Direction == currentCoord.Direction {
						shouldAppend = false
					}
				}
				if shouldAppend {
					previous[edge.CoordWithDirection] = append(previous[edge.CoordWithDirection], currentCoord)
				}
				heap.Push(&pq, &utils.PqItem{Value: edge.CoordWithDirection, Priority: newDistance})
			}
		}
	}
	path := make([]CoordWithDirection, 0)
	uniquePaths := make(map[CoordWithDirection]bool)
	doDfsOnPrevious(previous, starting, ending, path, uniquePaths)
	for c := range uniquePaths {
		grid[c.Coord.Row][c.Coord.Col] = "O"
	}
	count := 0
	ROWS, COLS := len(grid), len(grid[0])
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if grid[i][j] == "O" {
				count += 1
			}
		}
	}
	return count
}

func generateDay16AdjacencyList(grid [][]string) map[CoordWithDirection][]Day16AdjacencyNode {
	adjList := make(map[CoordWithDirection][]Day16AdjacencyNode)
	ROWS, COLS := len(grid), len(grid[0])
	directionToDirectionsMap := map[string][]string{
		"up":    []string{"up", "right", "left"},
		"right": []string{"right", "up", "down"},
		"down":  []string{"down", "left", "right"},
		"left":  []string{"left", "up", "down"},
	}
	directionToCoordsMap := map[string]utils.Coord{
		"up":    utils.Grid4Directions[0],
		"down":  utils.Grid4Directions[1],
		"left":  utils.Grid4Directions[2],
		"right": utils.Grid4Directions[3],
	}
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if grid[i][j] != "#" {
				for _, direction := range []string{"up", "down", "left", "right"} {
					currCoord := CoordWithDirection{
						Coord:     utils.Coord{Row: i, Col: j},
						Direction: direction,
					}
					adjList[currCoord] = make([]Day16AdjacencyNode, 0)
					for _, nextDirection := range directionToDirectionsMap[direction] {
						d := directionToCoordsMap[nextDirection]
						nx, ny := i+d.Row, j+d.Col
						if nx >= 0 && nx < ROWS && ny >= 0 && ny < COLS && grid[nx][ny] != "#" {
							adjacencyNode := Day16AdjacencyNode{
								CoordWithDirection: CoordWithDirection{
									Coord:     utils.Coord{Row: nx, Col: ny},
									Direction: nextDirection,
								},
							}
							if nextDirection != direction {
								adjacencyNode.Distance = 1001
							} else {
								adjacencyNode.Distance = 1
							}
							adjList[currCoord] = append(adjList[currCoord], adjacencyNode)
						}
					}
				}
			}
		}
	}
	return adjList
}

func doDfsOnPrevious(
	m map[CoordWithDirection][]CoordWithDirection,
	start CoordWithDirection,
	curr CoordWithDirection,
	path []CoordWithDirection,
	uniquePaths map[CoordWithDirection]bool,
) {
	path = append(path, curr)
	if curr.Coord == start.Coord && curr.Direction == start.Direction {
		// Print out different path lengths to find out the shortest path length to filter for.
		if len(path) == 501 {
			for _, p := range path {
				uniquePaths[p] = true
			}
		}
		return
	}
	for _, v := range m[curr] {
		newPath := make([]CoordWithDirection, 0)
		newPath = append(newPath, path...)
		doDfsOnPrevious(m, start, v, newPath, uniquePaths)
	}
}
