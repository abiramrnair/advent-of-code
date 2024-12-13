package solutions

import (
	"aoc/utils"
	"fmt"
)

func Day8_2024_Part1() {
	defer utils.CodeTimer()()
	INPUT_PATH := "./inputs/Year2024Day8.txt"
	input, err := utils.GetInputAs2DArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	ROWS, COLS := len(input), len(input[0])
	coords := make([]utils.Coord, 0)
	uniquePos := make(map[utils.Coord]bool)
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if input[i][j] != "." {
				coords = append(coords, utils.Coord{
					Row: i,
					Col: j,
				})
			}
		}
	}
	for i := 0; i < len(coords); i++ {
		nodeone := coords[i]
		antennaone := input[nodeone.Row][nodeone.Col]
		for j := 0; j < len(coords); j++ {
			nodetwo := coords[j]
			antennatwo := input[nodetwo.Row][nodetwo.Col]
			if i != j && antennaone == antennatwo {
				xdiff := nodeone.Row - nodetwo.Row
				ydiff := nodeone.Col - nodetwo.Col
				possibles := []utils.Coord{
					{
						Row: nodeone.Row + xdiff,
						Col: nodeone.Col + ydiff,
					},
					{
						Row: nodeone.Row - xdiff,
						Col: nodeone.Col - ydiff,
					},
					{
						Row: nodetwo.Row + xdiff,
						Col: nodetwo.Col + ydiff,
					},
					{
						Row: nodetwo.Row - xdiff,
						Col: nodetwo.Col - ydiff,
					},
				}
				for _, p := range possibles {
					if (p.Row == nodeone.Row && p.Col == nodeone.Col) ||
						(p.Row == nodetwo.Row && p.Col == nodetwo.Col) {
						continue
					}
					if p.Row >= 0 && p.Row < ROWS && p.Col >= 0 && p.Col < COLS {
						coord := utils.Coord{Row: p.Row, Col: p.Col}
						_, ok := uniquePos[coord]
						if !ok {
							uniquePos[coord] = true
						}
					}
				}
			}
		}
	}
	fmt.Println(len(uniquePos))
}

func Day8_2024_Part2() {
	defer utils.CodeTimer()()
	INPUT_PATH := "./inputs/Year2024Day8.txt"
	input, err := utils.GetInputAs2DArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	ROWS, COLS := len(input), len(input[0])
	coords := make([]utils.Coord, 0)
	uniquePos := make(map[utils.Coord]bool)
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if input[i][j] != "." {
				coords = append(coords, utils.Coord{
					Row: i,
					Col: j,
				})
			}
		}
	}
	for i := 0; i < len(coords); i++ {
		nodeone := coords[i]
		antennaone := input[nodeone.Row][nodeone.Col]
		for j := 0; j < len(coords); j++ {
			nodetwo := coords[j]
			antennatwo := input[nodetwo.Row][nodetwo.Col]
			if i != j && antennaone == antennatwo {
				xdiff := nodeone.Row - nodetwo.Row
				ydiff := nodeone.Col - nodetwo.Col
				var possibles []utils.Coord
				i := 1
				for {
					nr := nodeone.Row + (i * xdiff)
					nc := nodeone.Col + (i * ydiff)
					if nr < 0 || nr > ROWS-1 || nc < 0 || nc > COLS-1 {
						break
					}
					possibles = append(possibles, utils.Coord{Row: nr, Col: nc})
					i += 1
				}
				i = 1
				for {
					nr := nodeone.Row - (i * xdiff)
					nc := nodeone.Col - (i * ydiff)
					if nr < 0 || nr > ROWS-1 || nc < 0 || nc > COLS-1 {
						break
					}
					possibles = append(possibles, utils.Coord{Row: nr, Col: nc})
					i += 1
				}
				for _, p := range possibles {
					coord := utils.Coord{Row: p.Row, Col: p.Col}
					_, ok := uniquePos[coord]
					if !ok {
						uniquePos[coord] = true
					}
				}
			}
		}
	}
	fmt.Println(len(uniquePos))
}
