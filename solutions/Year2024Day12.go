package solutions

import (
	"aoc/utils"
	"fmt"
	"sort"
)

func Day12_2024_Part1() {
	defer utils.CodeTimer()()
	INPUT_PATH := "./inputs/Year2024Day12.txt"
	input, err := utils.GetInputAs2DArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	ROWS, COLS := len(input), len(input[0])
	price := 0
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			letter := input[r][c]
			if utils.StringIsUppercase(letter) {
				price += gardenFloodFill(input, r, c)
			}
		}
	}
	fmt.Println("Answer:", price)
}

func Day12_2024_Part2() {
	defer utils.CodeTimer()()
	INPUT_PATH := "./inputs/Year2024Day12.txt"
	input, err := utils.GetInputAs2DArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	ROWS, COLS := len(input), len(input[0])
	price := 0
	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			letter := input[r][c]
			if utils.StringIsUppercase(letter) {
				price += gardenFloodFill2(input, r, c)
			}
		}
	}
	fmt.Println("Answer:", price)
}

func gardenFloodFill(grid [][]string, i int, j int) int {
	ROWS, COLS := len(grid), len(grid[0])
	area := 0
	perimeter := 0
	var q utils.Queue
	coord := utils.Coord{Row: i, Col: j}
	charToCompare := grid[i][j]
	visited := make(map[utils.Coord]bool)
	q.Enqueue(coord)

	for q.Size() > 0 {
		length := q.Size()
		i := 0
		for i < length {
			element := q.Dequeue().(utils.Coord)
			_, ok := visited[element]
			if !ok {
				visited[element] = true
				perimeter += 4
			} else {
				i += 1
				continue
			}
			area += 1
			r, c := element.Row, element.Col
			grid[r][c] = utils.ConvertStringToLowercase(charToCompare)
			for _, dir := range utils.Grid4Directions {
				dx, dy := dir.Row, dir.Col
				nx, ny := r+dx, c+dy
				newCoord := utils.Coord{Row: nx, Col: ny}
				if nx >= 0 && nx < ROWS && ny >= 0 && ny < COLS && grid[nx][ny] == charToCompare {
					q.Enqueue(newCoord)
				}
				if nx >= 0 && nx < ROWS && ny >= 0 && ny < COLS && ((grid[nx][ny] == charToCompare) || (grid[nx][ny] == utils.ConvertStringToLowercase(charToCompare))) {
					perimeter -= 1
				}
			}
			i += 1
		}
	}
	return area * perimeter
}

func gardenFloodFill2(grid [][]string, i int, j int) int {
	ROWS, COLS := len(grid), len(grid[0])
	stringDirections := map[utils.Coord]string{
		utils.Grid4Directions[0]: "Down",
		utils.Grid4Directions[1]: "Up",
		utils.Grid4Directions[2]: "Right",
		utils.Grid4Directions[3]: "Left",
	}
	type NumAndDirection struct {
		Idx       int
		Type      string
		Direction string
	}
	numAndDirectionMap := make(map[NumAndDirection][]int)
	area := 0
	walls := 0
	var q utils.Queue
	coord := utils.Coord{Row: i, Col: j}
	charToCompare := grid[i][j]
	visited := make(map[utils.Coord]bool)
	q.Enqueue(coord)

	for q.Size() > 0 {
		length := q.Size()
		i := 0
		for i < length {
			element := q.Dequeue().(utils.Coord)
			_, ok := visited[element]
			if !ok {
				visited[element] = true
			} else {
				i += 1
				continue
			}
			area += 1
			r, c := element.Row, element.Col
			grid[r][c] = utils.ConvertStringToLowercase(charToCompare)
			for _, dir := range utils.Grid4Directions {
				dx, dy := dir.Row, dir.Col
				nx, ny := r+dx, c+dy
				newCoord := utils.Coord{Row: nx, Col: ny}
				if nx >= 0 && nx < ROWS && ny >= 0 && ny < COLS && grid[nx][ny] == charToCompare {
					q.Enqueue(newCoord)
				}
				if (nx >= 0 && nx < ROWS && ny >= 0 && ny < COLS &&
					grid[nx][ny] != charToCompare &&
					grid[nx][ny] != utils.ConvertStringToLowercase(charToCompare)) ||
					(nx < 0 || nx > ROWS-1 || ny < 0 || ny > COLS-1) {
					stringDirection := stringDirections[dir]
					if stringDirection == "Right" || stringDirection == "Left" {
						nd := NumAndDirection{Idx: c, Direction: stringDirection, Type: "Column"}
						_, ok := numAndDirectionMap[nd]
						if !ok {
							numAndDirectionMap[nd] = make([]int, 0)
						}
						numAndDirectionMap[nd] = append(numAndDirectionMap[nd], r)
					}
					if stringDirection == "Up" || stringDirection == "Down" {
						nd := NumAndDirection{Idx: r, Direction: stringDirection, Type: "Row"}
						_, ok := numAndDirectionMap[nd]
						if !ok {
							numAndDirectionMap[nd] = make([]int, 0)
						}
						numAndDirectionMap[nd] = append(numAndDirectionMap[nd], c)
					}
				}
			}
			i += 1
		}
	}
	for _, v := range numAndDirectionMap {
		sort.Ints(v)
		for i := 0; i < len(v); i++ {
			if i == 0 {
				walls += 1
			} else if v[i]-v[i-1] > 1 {
				walls += 1
			}
		}
	}
	return area * walls
}
