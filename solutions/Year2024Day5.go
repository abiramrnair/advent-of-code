package solutions

import (
	"aoc/utils"
	"fmt"
)

func Day5_2024_Part1() {
	path := "./inputs/Year2024Day5.txt"
	input, err := utils.GetInputAsArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	backwardNums := make(map[int][]int)
	updates := make([][]int, 0)
	shouldUpdate := false
	for _, line := range input {
		if line == "" {
			shouldUpdate = true
			continue
		}
		if shouldUpdate {
			updateNums := utils.StringSplitByChar(line, ",")
			update := make([]int, 0)
			for _, num := range updateNums {
				update = append(update, utils.ConvertStringToInt(num))
			}
			updates = append(updates, update)
		} else {
			fb := utils.StringSplitByChar(line, "|")
			backward, forward := utils.ConvertStringToInt(fb[0]), utils.ConvertStringToInt(fb[1])
			_, ok := backwardNums[forward]
			if !ok {
				backwardNums[forward] = make([]int, 0)
			}
			backwardNums[forward] = append(backwardNums[forward], backward)
		}
	}
	sum := 0
	for _, update := range updates {
		validUpdate := isUpdateValid(update, backwardNums)
		if validUpdate {
			midIdx := 0 + (len(update)-0)/2
			sum += update[midIdx]
		}
	}
	fmt.Println(sum)
}

func Day5_2024_Part2() {
	path := "./inputs/Year2024Day5.txt"
	input, err := utils.GetInputAsArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	backwardNums := make(map[int][]int)
	updates := make([][]int, 0)
	shouldUpdate := false
	for _, line := range input {
		if line == "" {
			shouldUpdate = true
			continue
		}
		if shouldUpdate {
			updateNums := utils.StringSplitByChar(line, ",")
			update := make([]int, 0)
			for _, num := range updateNums {
				update = append(update, utils.ConvertStringToInt(num))
			}
			updates = append(updates, update)
		} else {
			fb := utils.StringSplitByChar(line, "|")
			backward, forward := utils.ConvertStringToInt(fb[0]), utils.ConvertStringToInt(fb[1])
			_, ok := backwardNums[forward]
			if !ok {
				backwardNums[forward] = make([]int, 0)
			}
			backwardNums[forward] = append(backwardNums[forward], backward)
		}
	}
	sum := 0
	for z := range updates {
		isValid := isUpdateValid(updates[z], backwardNums)
		if !isValid {
			for i := len(updates[z]) - 1; i >= 0; i-- {
				for j := i - 1; j >= 0; j-- {
					if !isValidPosition(updates[z][i], updates[z][j], backwardNums) {
						tmp := updates[z][j]
						updates[z][j] = updates[z][i]
						updates[z][i] = tmp
					}
				}
			}
			midIdx := 0 + (len(updates[z])-0)/2
			sum += updates[z][midIdx]
		}
	}
	fmt.Println(sum)
}

func isUpdateValid(update []int, backwardNums map[int][]int) bool {
	for i := len(update) - 1; i >= 0; i-- {
		currPage := update[i]
		beforePages := backwardNums[currPage]
		for j := i - 1; j >= 0; j-- {
			isFound := false
			prevPage := update[j]
			for z := 0; z < len(beforePages); z++ {
				if prevPage == beforePages[z] {
					isFound = true
				}
			}
			if !isFound {
				return false
			}
		}
	}
	return true
}

func isValidPosition(currNum int, prevNum int, backwardNums map[int][]int) bool {
	isValid := false
	for _, num := range backwardNums[currNum] {
		if num == prevNum {
			isValid = true
		}
	}
	return isValid
}
