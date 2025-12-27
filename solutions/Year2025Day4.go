package solutions

import (
	"aoc/utils"
	"fmt"
)

func Day4_2025_Part1() {
	defer utils.CodeTimer()()
	path := "./inputs/Year2025Day4.txt"
	input, err := utils.GetInputAs2DArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	ROWS, COLS := len(input), len(input[0])
	res := 0
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			char := input[i][j]
			ctr := 0
			if char == "@" {
				if j < COLS-1 && input[i][j+1] == "@" {
					ctr += 1
				}
				if j > 0 && input[i][j-1] == "@" {
					ctr += 1
				}
				if i > 0 && input[i-1][j] == "@" {
					ctr += 1
				}
				if i < ROWS-1 && input[i+1][j] == "@" {
					ctr += 1
				}
				if i < ROWS-1 && j < COLS-1 && input[i+1][j+1] == "@" {
					ctr += 1
				}
				if i < ROWS-1 && j > 0 && input[i+1][j-1] == "@" {
					ctr += 1
				}
				if i > 0 && j < COLS-1 && input[i-1][j+1] == "@" {
					ctr += 1
				}
				if i > 0 && j > 0 && input[i-1][j-1] == "@" {
					ctr += 1
				}
				if ctr < 4 {
					res += 1
				}
			}
		}
	}
	fmt.Println(res)
}

func Day4_2025_Part2() {
	defer utils.CodeTimer()()
	path := "./inputs/Year2025Day4.txt"
	input, err := utils.GetInputAs2DArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	ROWS, COLS := len(input), len(input[0])
	res := 0
	shouldLoop := true
	for shouldLoop {
		found := 0
		toClear := make([][2]int, 0)
		for i := 0; i < ROWS; i++ {
			for j := 0; j < COLS; j++ {
				char := input[i][j]
				ctr := 0
				if char == "@" {
					if j < COLS-1 && input[i][j+1] == "@" {
						ctr += 1
					}
					if j > 0 && input[i][j-1] == "@" {
						ctr += 1
					}
					if i > 0 && input[i-1][j] == "@" {
						ctr += 1
					}
					if i < ROWS-1 && input[i+1][j] == "@" {
						ctr += 1
					}
					if i < ROWS-1 && j < COLS-1 && input[i+1][j+1] == "@" {
						ctr += 1
					}
					if i < ROWS-1 && j > 0 && input[i+1][j-1] == "@" {
						ctr += 1
					}
					if i > 0 && j < COLS-1 && input[i-1][j+1] == "@" {
						ctr += 1
					}
					if i > 0 && j > 0 && input[i-1][j-1] == "@" {
						ctr += 1
					}
					if ctr < 4 {
						found += 1
						res += 1
						toClear = append(toClear, [2]int{i, j})
					}
				}
			}
		}
		if found == 0 {
			shouldLoop = false
		}
		for _, coord := range toClear {
			i, j := coord[0], coord[1]
			input[i][j] = "."
		}
	}
	fmt.Println(res)
}
