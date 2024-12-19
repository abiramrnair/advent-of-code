package solutions

import (
	"aoc/utils"
	"fmt"
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
	answer := djikstraDay16(adjacencyList)
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

func djikstraDay16(adjacencyList map[CoordWithDirection][]Day16AdjacencyNode) int {
	fmt.Println(adjacencyList)
	return 0
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
								adjacencyNode.Distance = 1000
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
