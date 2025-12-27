package solutions

import (
	"aoc/utils"
	"fmt"
	"sort"
	"strings"
)

func Day5_2025_Part1() {
	defer utils.CodeTimer()()
	path := "./inputs/Year2025Day5.txt"
	input, err := utils.GetInputAsArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	ranges := make([]string, 0)
	ids := make([]string, 0)
	getIds := false
	res := 0
	for _, line := range input {
		if line == "" {
			getIds = true
			continue
		}
		if getIds {
			ids = append(ids, line)
		} else {
			ranges = append(ranges, line)
		}
	}
	for _, id := range ids {
		isValid := false
		num := utils.ConvertStringToInt(id)
		for _, item := range ranges {
			r := strings.Split(item, "-")
			start, end := utils.ConvertStringToInt(r[0]), utils.ConvertStringToInt(r[1])
			if num >= start && num <= end {
				isValid = true
			}
		}
		if isValid {
			res += 1
		}
	}
	fmt.Println(res)
}

func Day5_2025_Part2() {
	defer utils.CodeTimer()()
	path := "./inputs/Year2025Day5.txt"
	input, err := utils.GetInputAsArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	ranges := make([][2]int, 0)
	for _, line := range input {
		if line == "" {
			break
		}
		str := strings.Split(line, "-")
		start, end := utils.ConvertStringToInt(str[0]), utils.ConvertStringToInt(str[1])
		ranges = append(ranges, [2]int{start, end})
	}
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	i := 0
	j := 1
	newRanges := make([][2]int, 0, len(ranges))
	for j < len(ranges) {
		firstInterval := ranges[i]
		secondInterval := ranges[j]
		if firstInterval[1] >= secondInterval[0] {
			firstInterval[1] = utils.MaxNumber(firstInterval[1], secondInterval[1])
			ranges[i] = firstInterval
		} else {
			newRanges = append(newRanges, ranges[i])
			i = j
		}
		j += 1
	}
	newRanges = append(newRanges, ranges[i])
	res := 0
	for _, item := range newRanges {
		start, end := item[0], item[1]
		res += (end - start) + 1
	}
	fmt.Println(res)
}
