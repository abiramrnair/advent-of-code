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
	ROWS,COLS := len(input),len(input[0])
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
	input, err := utils.GetInputAsArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(input)
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
	pq := make(utils.PriorityQueue,0)
	heap.Init(&pq)
	heap.Push(&pq, &utils.PqItem{Value: starting, Priority: 0})
	for pq.Len() > 0 {
		currentNode := heap.Pop(&pq).(*utils.PqItem)
		currentCoord := currentNode.Value.(CoordWithDirection)
		currentDistance := currentNode.Priority
		visited[currentCoord] = true
		for _,edge := range adjacencyList[currentCoord] {
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
	for k,v := range distances {
		if k.Row == ending.Row && k.Col == ending.Col {
			if v < minDistance {
				minDistance = v
			}
		}
	}
	return minDistance
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
