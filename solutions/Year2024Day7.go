package solutions

import (
	"aoc/utils"
	"fmt"
)

func Day7_2024_Part1() {
	defer utils.CodeTimer()()
	INPUT_PATH := "./inputs/Year2024Day7.txt"
	input, err := utils.GetInputAsArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	res := 0
	for _, line := range input {
		lineArgs := utils.StringSplitByChar(line, ":")
		total, nums := lineArgs[0], utils.StringSplitByChar(utils.StringTrimSpaces(lineArgs[1]), " ")
		var possibilities [][]string
		getPossibleNumsCombos(nums, []string{}, &possibilities, total, 0)
		if len(possibilities) > 0 {
			res += utils.ConvertStringToInt(total)
		}
	}
	fmt.Println(res)
}

func Day7_2024_Part2() {
	defer utils.CodeTimer()()
	INPUT_PATH := "./inputs/Year2024Day7.txt"
	input, err := utils.GetInputAsArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	res := 0
	for _, line := range input {
		lineArgs := utils.StringSplitByChar(line, ":")
		total, nums := lineArgs[0], utils.StringSplitByChar(utils.StringTrimSpaces(lineArgs[1]), " ")
		var possibilities [][]string
		getPossibleNumsCombos2(nums, []string{}, &possibilities, total, 0)
		if len(possibilities) > 0 {
			res += utils.ConvertStringToInt(total)
		}
	}
	fmt.Println(res)
}

func getPossibleNumsCombos(nums []string, path []string, possible *[][]string, total string, index int) {
	if index == len(nums) {
		if checkValidPath(total, path) {
			*possible = append(*possible, path)
		}
		return
	}
	path = append(path, nums[index])
	path = append(path, "+")
	getPossibleNumsCombos(nums, path, possible, total, index+1)
	path = path[:len(path)-1]
	path = append(path, "*")
	getPossibleNumsCombos(nums, path, possible, total, index+1)
}

func getPossibleNumsCombos2(nums []string, path []string, possible *[][]string, total string, index int) {
	if index == len(nums) {
		if checkValidPath2(total, path, false) {
			*possible = append(*possible, path)
		}
		return
	}
	if !checkValidPath2(total, path, true) {
		return
	}
	path = append(path, nums[index])
	path = append(path, "+")
	getPossibleNumsCombos2(nums, path, possible, total, index+1)
	path = path[:len(path)-1]
	path = append(path, "*")
	getPossibleNumsCombos2(nums, path, possible, total, index+1)
	path = path[:len(path)-1]
	path = append(path, "||")
	getPossibleNumsCombos2(nums, path, possible, total, index+1)
}

func checkValidPath(total string, path []string) bool {
	res := 0
	for i, char := range path {
		if i == 0 {
			res = utils.ConvertStringToInt(char)
		} else if char == "+" || char == "*" {
			continue
		} else {
			prevChar := path[i-1]
			if prevChar == "+" {
				res += utils.ConvertStringToInt(char)
			} else if prevChar == "*" {
				res *= utils.ConvertStringToInt(char)
			}
		}
	}
	return utils.ConvertStringToInt(total) == res
}

func checkValidPath2(total string, path []string, trim bool) bool {
	res := ""
	for i, char := range path {
		if i == 0 {
			res = char
		} else if char == "+" || char == "*" {
			continue
		} else {
			prevChar := path[i-1]
			if prevChar == "+" {
				intres := utils.ConvertStringToInt(res)
				intres += utils.ConvertStringToInt(char)
				res = utils.ConvertIntToString(intres)
			} else if prevChar == "*" {
				intres := utils.ConvertStringToInt(res)
				intres *= utils.ConvertStringToInt(char)
				res = utils.ConvertIntToString(intres)
			} else if prevChar == "||" {
				res += char
			}
		}
	}
	if trim {
		return !(utils.ConvertStringToInt(res) > utils.ConvertStringToInt(total))
	} else {
		return utils.ConvertStringToInt(total) == utils.ConvertStringToInt(res)
	}
}
