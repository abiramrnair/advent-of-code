package solutions

import (
	"aoc/utils"
	"fmt"
)

func Day10_2024_Part1() {
	defer utils.CodeTimer()()
	INPUT_PATH := "./inputs/Year2024Day10.txt"
	input, err := utils.GetInputAs2DArrayOfInts(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	answer := 0
	ROWS, COLS := len(input), len(input[0])
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			square := input[i][j]
			if square == 0 {
				visited := make(map[utils.Coord]bool)
				answer += findTrailheadSum(input, i, j, visited)
			}
		}
	}
	fmt.Println("Answer:", answer)
}

func Day10_2024_Part2() {
	defer utils.CodeTimer()()
	INPUT_PATH := "./inputs/Year2024Day10.txt"
	input, err := utils.GetInputAs2DArrayOfInts(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	answer := 0
	ROWS, COLS := len(input), len(input[0])
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			square := input[i][j]
			if square == 0 {
				answer += findTrailheadSum2(input, i, j)
			}
		}
	}
	fmt.Println("Answer:", answer)
}

func findTrailheadSum(arr [][]int, i int, j int, visited map[utils.Coord]bool) int {
	ROWS, COLS := len(arr), len(arr[0])
	val := arr[i][j]
	if val == 9 {
		return 1
	}
	res := 0
	for _, dir := range utils.Grid4Directions {
		dx, dy := dir.Row, dir.Col
		nx, ny := i+dx, j+dy
		newCoord := utils.Coord{Row: nx, Col: ny}
		if nx >= 0 && nx < ROWS && ny >= 0 && ny < COLS {
			nxt := arr[nx][ny]
			_, ok := visited[newCoord]
			if nxt == (val+1) && !ok {
				visited[newCoord] = true
				res += findTrailheadSum(arr, nx, ny, visited)
			}
		}
	}
	return res
}

func findTrailheadSum2(arr [][]int, i int, j int) int {
	ROWS, COLS := len(arr), len(arr[0])
	val := arr[i][j]
	if val == 9 {
		return 1
	}
	res := 0
	for _, dir := range utils.Grid4Directions {
		dx, dy := dir.Row, dir.Col
		nx, ny := i+dx, j+dy
		if nx >= 0 && nx < ROWS && ny >= 0 && ny < COLS {
			nxt := arr[nx][ny]
			if nxt == (val + 1) {
				res += findTrailheadSum2(arr, nx, ny)
			}
		}
	}
	return res
}
