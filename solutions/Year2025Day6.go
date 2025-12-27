package solutions

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func Day6_2025_Part1() {
	defer utils.CodeTimer()()
	path := "./inputs/Year2025Day6.txt"
	input, err := utils.GetInputAsArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	numsArray := make([][]int, 0)
	opsArray := make([]string, 0)
	for i, line := range input {
		if i == len(input)-1 {
			strArr := strings.Split(line, " ")
			for _, char := range strArr {
				if char == "*" || char == "+" {
					opsArray = append(opsArray, char)
				}
			}
		}
		nums := make([]int, 0)
		strArr := strings.Split(line, " ")
		for _, char := range strArr {
			toAdd := utils.ConvertStringToInt(char)
			if toAdd != 0 {
				nums = append(nums, toAdd)
			}
		}
		if len(nums) > 0 {
			numsArray = append(numsArray, nums)
		}
	}
	ROWS, COLS := len(numsArray), len(numsArray[0])
	total := 0
	for j := 0; j < COLS; j++ {
		operation := opsArray[j]
		res := 0
		if operation == "*" {
			res = 1
		}
		for i := 0; i < ROWS; i++ {
			if operation == "*" {
				res *= numsArray[i][j]
			} else if operation == "+" {
				res += numsArray[i][j]
			}
		}
		total += res
	}
	fmt.Println(total)
}

func Day6_2025_Part2() {
	defer utils.CodeTimer()()
	path := "./inputs/Year2025Day6.txt"
	input, err := utils.GetInputAs2DArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	ROWS, COLS := len(input), len(input[0])
	for j := 0; j < COLS; j++ {
		if input[0][j] == " " {
			i := 0
			isZeros := true
			for i < ROWS-1 {
				if input[i][j] != " " {
					isZeros = false
				}
				i += 1
			}
			if isZeros {
				i := 0
				for i < ROWS-1 {
					input[i][j] = "#"
					i += 1
				}
			}
		}
	}
	numsArray := make([][][]string, 0)
	opsArray := make([]string, 0)
	for i := ROWS - 1; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if input[i][j] != " " {
				opsArray = append(opsArray, input[i][j])
			}
		}
	}
	for j, line := range input {
		if j == len(input)-1 {
			continue
		}
		toAdd := make([]string, 0)
		strs := make([][]string, 0)
		for i, c := range line {
			if c == "#" {
				strs = append(strs, toAdd)
				toAdd = make([]string, 0)
			} else {
				toAdd = append(toAdd, c)
				if i == len(line)-1 {
					strs = append(strs, toAdd)
				}
			}
		}
		numsArray = append(numsArray, strs)
	}
	ROWS, COLS = len(numsArray), len(numsArray[0])
	ans := 0
	for j := COLS - 1; j >= 0; j-- {
		columnLength := len(numsArray[0][j])
		opToUse := opsArray[j]
		res := 0
		if opToUse == "*" {
			res = 1
		}
		for z := columnLength - 1; z >= 0; z-- {
			numToOp := ""
			for i := 0; i < ROWS; i++ {
				if numsArray[i][j][z] != " " {
					numToOp += numsArray[i][j][z]
				}
			}
			if opToUse == "+" {
				res += utils.ConvertStringToInt(numToOp)
			} else if opToUse == "*" {
				res *= utils.ConvertStringToInt(numToOp)
			}
		}
		ans += res
	}
	fmt.Println(ans)
}
