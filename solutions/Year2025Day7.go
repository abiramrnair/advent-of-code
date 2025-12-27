package solutions

import (
	"aoc/utils"
	"fmt"
)

func Day7_2025_Part1() {
	defer utils.CodeTimer()()
	path := "./inputs/Year2025Day7.txt"
	input, err := utils.GetInputAs2DArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	// Find starting point
	ROWS, COLS := len(input), len(input[0])
	res := 0
	visited := map[utils.Coord]bool{}
	var findSplitter func(i, j int)
	findSplitter = func(i int, j int) {
		if input[i][j] == "^" {
			pos := utils.Coord{Row: i, Col: j}
			_, ok := visited[pos]
			if !ok {
				res += 1
				visited[pos] = true
			}
			if j+1 < COLS {
				findSplitter(i, j+1)
			}
			if j-1 >= 0 {
				findSplitter(i, j-1)
			}
		} else {
			input[i][j] = "|"
			pos := utils.Coord{Row: i, Col: j}
			visited[pos] = true
			if i+1 < ROWS {
				newPos := utils.Coord{Row: i + 1, Col: j}
				_, ok := visited[newPos]
				if !ok {
					findSplitter(i+1, j)
				}
			}
		}
	}
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if input[i][j] == "S" {
				findSplitter(i+1, j)
				break
			}
		}
	}
	fmt.Println(res)
}

func Day7_2025_Part2() {
	defer utils.CodeTimer()()
	path := "./inputs/Year2025Day7.txt"
	input, err := utils.GetInputAs2DArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	ROWS, COLS := len(input), len(input[0])
	res := 0
	memo := map[utils.Coord]int{}
	var findSplitter func(i int, j int, steps int) int
	findSplitter = func(i int, j int, steps int) int {
		for input[i][j] == "." && i < ROWS-1 {
			i += 1
		}
		if input[i][j] == "^" {
			coord := utils.Coord{Row: i, Col: j}
			_, ok := memo[coord]
			if ok {
				return memo[coord]
			}
			steps += findSplitter(i, j+1, steps) + findSplitter(i, j-1, steps)
			memo[coord] = steps
			return memo[coord]
		}
		return 1
	}
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if input[i][j] == "S" {
				res = findSplitter(i+1, j, 0)
				break
			}
		}
	}
	fmt.Println(res)
}
