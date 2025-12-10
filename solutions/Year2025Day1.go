package solutions

import (
	"aoc/utils"
	"fmt"
)

func Day1_2025_Part1() {
	path := "./inputs/Year2025Day1.txt"
	input, err := utils.GetInputAsArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	zeroRotations := 0
	currentPosition := 50
	for _, rotation := range input {
		direction := utils.StringSplitByChar(rotation, "")[0]
		amount := utils.ConvertStringToInt(utils.StringJoinByChar(utils.StringSplitByChar(rotation, "")[1:], ""))
		if direction == "R" {
			currentPosition += amount
		} else if direction == "L" {
			currentPosition -= amount
		}
		if currentPosition%100 == 0 {
			zeroRotations += 1
		}
	}
	fmt.Println(zeroRotations)
}

func Day1_2025_Part2() {
	path := "./inputs/Year2025Day1.txt"
	input, err := utils.GetInputAsArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	zeroRotations := 0
	currentPosition := 50
	for _, rotation := range input {
		direction := utils.StringSplitByChar(rotation, "")[0]
		amount := utils.ConvertStringToInt(utils.StringJoinByChar(utils.StringSplitByChar(rotation, "")[1:], ""))
		if direction == "R" {
			for i := 0; i < amount; i++ {
				currentPosition += 1
				if currentPosition == 100 {
					currentPosition = 0
				}
				if currentPosition == 0 {
					zeroRotations += 1
				}
			}
		} else if direction == "L" {
			for i := amount; i > 0; i-- {
				currentPosition -= 1
				if currentPosition == -1 {
					currentPosition = 99
				}
				if currentPosition == 0 {
					zeroRotations += 1
				}
			}
		}
	}
	fmt.Println(zeroRotations)
}
