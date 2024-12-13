package solutions

import (
	"aoc/utils"
	"fmt"
	"sync"
)

func Day6_2024_Part1() {
	INPUT_PATH := "./inputs/Year2024Day6.txt"
	GUARD := "^"
	OBSTACLE := "#"
	input, err := utils.GetInputAs2DArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	ROWS, COLS := len(input), len(input[0])
	var startPos []int
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			char := input[i][j]
			if char == GUARD {
				startPos = []int{i, j}
			}
		}
	}
	i, j := startPos[0], startPos[1]
	currDir := []int{-1, 0}
	distinctPos := make(map[utils.Coord]bool)
	for i >= 0 && i < ROWS && j >= 0 && j < COLS {
		coord := utils.Coord{Row: i, Col: j}
		_, ok := distinctPos[coord]
		if !ok {
			distinctPos[coord] = true
		}
		dx, dy := currDir[0], currDir[1]
		nx, ny := i+dx, j+dy
		if nx < 0 || nx > ROWS-1 || ny < 0 || ny > COLS-1 {
			break
		}
		next := input[nx][ny]
		if next == OBSTACLE {
			currDir = getNewDirection(currDir)
			continue
		}
		i, j = nx, ny
	}
	fmt.Println(len(distinctPos))
}

func Day6_2024_Part2() {
	INPUT_PATH := "./inputs/Year2024Day6.txt"
	defer utils.CodeTimer()()
	input, err := utils.GetInputAs2DArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	ROWS, COLS := len(input), len(input[0])
	distinctPos := 0
	var wg sync.WaitGroup
	for x := 0; x < ROWS; x++ {
		for z := 0; z < COLS; z++ {
			wg.Add(1)
			go tryGrid(x, z, &distinctPos, &wg)
		}
	}
	wg.Wait()
	fmt.Println(distinctPos)
}

func getNewDirection(currDir []int) []int {
	dx, dy := currDir[0], currDir[1]
	if dx == 0 && dy == 1 {
		return []int{1, 0}
	} else if dx == 0 && dy == -1 {
		return []int{-1, 0}
	} else if dx == 1 && dy == 0 {
		return []int{0, -1}
	} else if dx == -1 && dy == 0 {
		return []int{0, 1}
	}
	return currDir
}

func tryGrid(x int, z int, total *int, wg *sync.WaitGroup) {
	INPUT_PATH := "./inputs/Year2024Day6.txt"
	GUARD := "^"
	OBSTACLE := "#"
	NEWOBSTACLE := "O"
	EMPTY := "."
	defer wg.Done()
	input, err := utils.GetInputAs2DArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	ROWS, COLS := len(input), len(input[0])
	var startPos []int
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			char := input[i][j]
			if char == GUARD {
				startPos = []int{i, j}
			}
		}
	}
	if input[x][z] == EMPTY {
		input[x][z] = NEWOBSTACLE
		distinctPos := make(map[utils.Coord]int)
		i, j := startPos[0], startPos[1]
		currDir := []int{-1, 0}
		for i >= 0 && i < ROWS && j >= 0 && j < COLS {
			coord := utils.Coord{Row: i, Col: j}
			_, ok := distinctPos[coord]
			if !ok {
				distinctPos[coord] = 0
			}
			if ok {
				distinctPos[coord] += 1
				if distinctPos[coord] > 4 {
					*total += 1
					break
				}
			}
			dx, dy := currDir[0], currDir[1]
			nx, ny := i+dx, j+dy
			if nx < 0 || nx > ROWS-1 || ny < 0 || ny > COLS-1 {
				break
			}
			next := input[nx][ny]
			if next == OBSTACLE || next == NEWOBSTACLE {
				currDir = getNewDirection(currDir)
				continue
			}
			i, j = nx, ny
		}
		input[x][z] = EMPTY
	}
}
