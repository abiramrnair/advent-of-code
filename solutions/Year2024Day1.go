package solutions

import (
	"aoc/utils"
	"fmt"
	"math"
	"sort"
)

func Day1_2024_Part1() {
	path := "./inputs/Year2024Day1.txt"
	input, err := utils.GetInputAsArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	var leftList []int
	var rightList []int
	for _, row := range input {
		nums := utils.StringSplitByChar(row, "   ")
		left, right := nums[0], nums[1]
		leftList = append(leftList, utils.ConvertStringToInt(left))
		rightList = append(rightList, utils.ConvertStringToInt(right))
		sort.Ints(leftList)
		sort.Ints(rightList)
	}
	total := 0
	for i, num := range leftList {
		total += int(math.Abs(float64(num - rightList[i])))
	}
	fmt.Println(total)
}

func Day1_2024_Part2() {
	path := "./inputs/Year2024Day1.txt"
	input, err := utils.GetInputAsArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	leftMap := make(map[int]int)
	var leftList []int
	var rightList []int
	for _, row := range input {
		nums := utils.StringSplitByChar(row, "   ")
		left, right := utils.ConvertStringToInt(nums[0]), utils.ConvertStringToInt(nums[1])
		_, ok := leftMap[left]
		if !ok {
			leftMap[left] = 0
		}
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}
	for _, num := range rightList {
		_, ok := leftMap[num]
		if ok {
			leftMap[num] += 1
		}
	}
	total := 0
	for _, num := range leftList {
		total += (num * leftMap[num])
	}
	fmt.Println(total)
}
