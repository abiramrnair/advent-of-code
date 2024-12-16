package solutions

import (
	"aoc/utils"
	"fmt"
)

type Box struct {
	Left utils.Coord
	Right utils.Coord
}

func Day15_2024_Part1() {
	defer utils.CodeTimer()()
	INPUT_PATH := "./inputs/Year2024Day15.txt"
	input, err := utils.GetInputAsArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	directions := make([]string,0)
	grid := make([][]string,0)
	marker := 0
	for i,line := range input {
		if line == "" {
			marker = i
			break
		}
		a := utils.StringSplitByChar(line,"")
		grid = append(grid, a)
	}
	for i,line := range input {
		if i >= marker {
			for _,char := range line {
				directions = append(directions, string(char))
			}
		}
	}
	for _,d := range directions {
		makeMove(grid,string(d))
	}
	ROWS,COLS := len(grid),len(grid[0])
	answer := 0
	for i := 0; i < ROWS; i++ {
		for j := 0 ; j < COLS; j++ {
			if grid[i][j] == "O" {
				answer += (i * 100) + j
			}
		}
	}
	fmt.Println("Answer:",answer)
}

func Day15_2024_Part2() {
	defer utils.CodeTimer()()
	INPUT_PATH := "./inputs/Year2024Day15.txt"
	input, err := utils.GetInputAsArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	directions := make([]string,0)
	grid := make([][]string,0)
	marker := 0
	for i,line := range input {
		if line == "" {
			marker = i
			break
		}
		a := utils.StringSplitByChar(line,"")
		n := make([]string,0)
		for _,char := range a {
			if char == "#" {
				n = append(n, []string{"#","#"}...)
			} else if char == "O" {
				n = append(n, []string{"[","]"}...)
			} else if char == "." {
				n = append(n, []string{".","."}...)
			} else if char == "@" {
				n = append(n, []string{"@","."}...)
			}
		}
		grid = append(grid, n)
	}
	for i,line := range input {
		if i >= marker {
			for _,char := range line {
				directions = append(directions, string(char))
			}
		}
	}
	for _,d := range directions {
		makeMove2(grid,string(d))
	}
	ROWS,COLS := len(grid),len(grid[0])
	answer := 0
	for i := 0; i < ROWS; i++ {
		for j := 0 ; j < COLS; j++ {
			if grid[i][j] == "[" {
				answer += (i * 100) + j
			}
		}
	}
	fmt.Println("Answer:",answer)
}

func makeMove(grid [][]string, direction string) {
	var currPos utils.Coord
	ROWS,COLS := len(grid),len(grid[0])
	for i := 0; i < ROWS; i++ {
		for j := 0 ; j < COLS; j++ {
			if grid[i][j] == "@" {
				currPos = utils.Coord{Row: i,Col: j}
				break
			}
		}
	}
	nsewMap := map[string]utils.Coord{
		"v": utils.Grid4Directions[0],
		"^": utils.Grid4Directions[1],
		">": utils.Grid4Directions[2],
		"<": utils.Grid4Directions[3],
	}
	nd := nsewMap[direction]
	dx,dy := nd.Row,nd.Col
	nx,ny := currPos.Row + dx,currPos.Col + dy
	if grid[nx][ny] == "." {
		grid[nx][ny] = "@"
		grid[currPos.Row][currPos.Col] = "."
	} else if grid[nx][ny] == "O" {
		coordsToMove := make([]utils.Coord, 0)
		i := 0
		for {
			cx,cy := currPos.Row + (i * dx),currPos.Col + (i *dy)
			if grid[cx][cy] != "O" && grid[cx][cy] != "@" {
				if grid[cx][cy] == "#" {
					coordsToMove = make([]utils.Coord,0)
				}
				break
			}
			coordsToMove = append(coordsToMove, utils.Coord{Row: cx,Col: cy})
			i += 1
		}
		for i := len(coordsToMove) - 1 ; i >= 0 ; i-- {
			toMove := coordsToMove[i]
			if grid[toMove.Row][toMove.Col] == "O" {
				grid[toMove.Row + dx][toMove.Col + dy] = "O"
				grid[toMove.Row][toMove.Col] = "."
			} else {
				grid[toMove.Row + dx][toMove.Col + dy] = "@"
				grid[toMove.Row][toMove.Col] = "."
			}
		}
	}
}

func makeMove2(grid [][]string, direction string) {
	var currPos utils.Coord
	ROWS,COLS := len(grid),len(grid[0])
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if grid[i][j] == "@" {
				currPos = utils.Coord{Row: i,Col: j}
				break
			}
		}
	}
	nsewMap := map[string]utils.Coord{
		"v": utils.Grid4Directions[0],
		"^": utils.Grid4Directions[1],
		">": utils.Grid4Directions[2],
		"<": utils.Grid4Directions[3],
	}
	nd := nsewMap[direction]
	dx,dy := nd.Row,nd.Col
	nx,ny := currPos.Row + dx,currPos.Col + dy
	if direction == ">" || direction == "<" {
		if grid[nx][ny] == "." {
			grid[nx][ny] = "@"
			grid[currPos.Row][currPos.Col] = "."
		} else if grid[nx][ny] == "]" {
			i := ny
			for grid[nx][i] != "." && grid[nx][i] != "#" {
				i -= 1
			}
			if grid[nx][i] == "." {
				for j := i; j <= ny; j++ {
					grid[nx][j] = grid[nx][j + 1]
					grid[nx][j + 1] = "."
				}
			}
		} else if grid[nx][ny] == "[" {
			i := ny
			for grid[nx][i] != "." && grid[nx][i] != "#" {
				i += 1
			}
			if grid[nx][i] == "." {
				for j := i; j >= ny; j-- {
					grid[nx][j] = grid[nx][j - 1]
					grid[nx][j - 1] = "."
				}
			}
		}
	} else if direction == "^" || direction == "v" {
		if grid[nx][ny] == "." {
			grid[nx][ny] = "@"
			grid[currPos.Row][currPos.Col] = "."
		} else if grid[nx][ny] == "]" || grid[nx][ny] == "[" {
			boxesMap := make(map[Box]bool)
			var z int 
			if direction == "^" {
				z = -1
			} else {
				z = 1
			}
			should := dfs(grid,nx,ny,z,boxesMap)
			if should {
				boxesArray := make([]Box,0)
				for k := range boxesMap {
					boxesArray = append(boxesArray, k)
				}
				counter := 0
				for counter < len(boxesArray) {
					for _,box := range boxesArray {
						leftBox := box.Left
						rightBox := box.Right
						if grid[leftBox.Row + z][leftBox.Col] == "." && grid[rightBox.Row + z][rightBox.Col] == "." {
							grid[leftBox.Row][leftBox.Col] = "."
							grid[rightBox.Row][rightBox.Col] = "."
							grid[leftBox.Row + z][leftBox.Col] = "["
							grid[rightBox.Row + z][rightBox.Col] = "]"
							counter += 1
						}
					}
				}
				grid[nx][ny] = "@"
				grid[currPos.Row][currPos.Col] = "."
			}
		}
	}
}

func dfs(grid [][]string, r int, c int, z int, boxes map[Box]bool) bool {
	char := grid[r][c]
	if char == "[" {
		leftBracketX,leftBracketY := r,c
		rightBracketX,rightBracketY := r,c+1
		boxes[Box{Left: utils.Coord{Row: leftBracketX,Col: leftBracketY},Right: utils.Coord{Row: rightBracketX,Col: rightBracketY}}] = true
		return dfs(grid,leftBracketX + z,leftBracketY,z,boxes) && dfs(grid,rightBracketX + z,rightBracketY,z,boxes)
	} else if char == "]" {
		rightBracketX,rightBracketY := r,c
		leftBracketX,leftBracketY := r,c-1
		boxes[Box{Left: utils.Coord{Row: leftBracketX,Col: leftBracketY},Right: utils.Coord{Row: rightBracketX,Col: rightBracketY}}] = true
		return dfs(grid,leftBracketX + z,leftBracketY,z,boxes) && dfs(grid,rightBracketX + z,rightBracketY,z,boxes)	
	} else if char == "#" {
		return false
	}
	return true
}