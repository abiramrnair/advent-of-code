package solutions

import (
	"aoc/utils"
	"fmt"
)

var count = 0

func Day4_2024_Part1() {
	path := "./inputs/Year2024Day4.txt"
	input, err := utils.GetInputAs2DArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	ROWS, COLS := len(input), len(input[0])
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if input[i][j] == "X" {
				determineWord(input, i, j)
			}
		}
	}
	fmt.Println(count)
}

func Day4_2024_Part2() {
	path := "./inputs/Year2024Day4.txt"
	input, err := utils.GetInputAs2DArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	ROWS, COLS := len(input), len(input[0])
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if input[i][j] == "A" {
				determineWord2(input, i, j)
			}
		}
	}
	fmt.Println(count)
}

func determineWord(grid [][]string, i int, j int) {
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}
	ROWS, COLS := len(grid), len(grid[0])
	for _, dir := range dirs {
		dx, dy := dir[0], dir[1]
		path := ""
		for z := 0; z < 4; z++ {
			nx, ny := i+(dx*z), j+(dy*z)
			if nx < 0 || nx > ROWS-1 || ny < 0 || ny > COLS-1 {
				break
			}
			path += grid[nx][ny]
			if path == "XMAS" {
				count += 1
			}
		}
	}
}

func determineWord2(grid [][]string, i int, j int) {
	dirs := [][]int{{1, 1}, {-1, 1}, {-1, -1}, {1, -1}}
	ROWS, COLS := len(grid), len(grid[0])
	mcoords, scoords := make([][]int, 0), make([][]int, 0)
	for _, dir := range dirs {
		dx, dy := dir[0], dir[1]
		nx, ny := i+dx, j+dy
		if nx < 0 || nx > ROWS-1 || ny < 0 || ny > COLS-1 {
			break
		}
		if grid[nx][ny] == "M" {
			mcoords = append(mcoords, []int{nx, ny})
		}
		if grid[nx][ny] == "S" {
			scoords = append(scoords, []int{nx, ny})
		}
	}
	if len(mcoords) == 2 && len(scoords) == 2 {
		if mcoords[0][0] != mcoords[1][0] && mcoords[0][1] != mcoords[1][1] {
			return
		}
		if scoords[0][0] != scoords[1][0] && scoords[0][1] != scoords[1][1] {
			return
		}
		count += 1
	}
}
