package solutions

import (
	"aoc/utils"
	"fmt"
)

func PartOne2023Day1() {
	path := "./inputs/Year2023Day1.txt"
	input, err := utils.GetInputAsArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	var seeds []string
	type Map struct {
		Start int
		End   int
		Step  int
	}
	var seedToSoilMap []Map
	var soilToFertilizerMap []Map
	var fertilizerToWaterMap []Map
	var waterToLightMap []Map
	var lightToTemperatureMap []Map
	var temperatureToHumidity []Map
	var humidityToLocation []Map
	idx := 0
	for {
		line := input[idx]
		if utils.StringHasChar(line, "seeds:") {
			seedStr := utils.StringSplitByChar(line, "seeds:")
			seeds = utils.StringSplitByChar(
				utils.StringTrimSpaces(seedStr[1]),
				" ",
			)
		} else if utils.StringHasChar(line, "seed-to-soil map:") {
			for {
				idx += 1
				line = input[idx]
				if line == "" {
					break
				}
				lineStr := utils.StringSplitByChar(line, " ")
				seedToSoilMap = append(seedToSoilMap, Map{
					Start: utils.ConvertStringToInt(lineStr[1]),
					End:   utils.ConvertStringToInt(lineStr[0]),
					Step:  utils.ConvertStringToInt(lineStr[2]),
				})
			}
		} else if utils.StringHasChar(line, "soil-to-fertilizer map:") {
			for {
				idx += 1
				line = input[idx]
				if line == "" {
					break
				}
				lineStr := utils.StringSplitByChar(line, " ")
				soilToFertilizerMap = append(soilToFertilizerMap, Map{
					Start: utils.ConvertStringToInt(lineStr[1]),
					End:   utils.ConvertStringToInt(lineStr[0]),
					Step:  utils.ConvertStringToInt(lineStr[2]),
				})
			}
		} else if utils.StringHasChar(line, "fertilizer-to-water map:") {
			for {
				idx += 1
				line = input[idx]
				if line == "" {
					break
				}
				lineStr := utils.StringSplitByChar(line, " ")
				fertilizerToWaterMap = append(fertilizerToWaterMap, Map{
					Start: utils.ConvertStringToInt(lineStr[1]),
					End:   utils.ConvertStringToInt(lineStr[0]),
					Step:  utils.ConvertStringToInt(lineStr[2]),
				})
			}
		} else if utils.StringHasChar(line, "water-to-light map:") {
			for {
				idx += 1
				line = input[idx]
				if line == "" {
					break
				}
				lineStr := utils.StringSplitByChar(line, " ")
				waterToLightMap = append(waterToLightMap, Map{
					Start: utils.ConvertStringToInt(lineStr[1]),
					End:   utils.ConvertStringToInt(lineStr[0]),
					Step:  utils.ConvertStringToInt(lineStr[2]),
				})
			}
		} else if utils.StringHasChar(line, "light-to-temperature map:") {
			for {
				idx += 1
				line = input[idx]
				if line == "" {
					break
				}
				lineStr := utils.StringSplitByChar(line, " ")
				lightToTemperatureMap = append(lightToTemperatureMap, Map{
					Start: utils.ConvertStringToInt(lineStr[1]),
					End:   utils.ConvertStringToInt(lineStr[0]),
					Step:  utils.ConvertStringToInt(lineStr[2]),
				})
			}
		} else if utils.StringHasChar(line, "temperature-to-humidity map:") {
			for {
				idx += 1
				line = input[idx]
				if line == "" {
					break
				}
				lineStr := utils.StringSplitByChar(line, " ")
				temperatureToHumidity = append(temperatureToHumidity, Map{
					Start: utils.ConvertStringToInt(lineStr[1]),
					End:   utils.ConvertStringToInt(lineStr[0]),
					Step:  utils.ConvertStringToInt(lineStr[2]),
				})
			}
		} else if utils.StringHasChar(line, "humidity-to-location map:") {
			for {
				idx += 1
				if idx > len(input)-1 {
					break
				}
				line = input[idx]
				if line == "" {
					break
				}
				lineStr := utils.StringSplitByChar(line, " ")
				humidityToLocation = append(humidityToLocation, Map{
					Start: utils.ConvertStringToInt(lineStr[1]),
					End:   utils.ConvertStringToInt(lineStr[0]),
					Step:  utils.ConvertStringToInt(lineStr[2]),
				})
			}
		}
		idx += 1
		if idx > len(input)-1 {
			break
		}
	}
	for _, s := range seeds {
		seed := utils.ConvertStringToInt(s)
		var soil int
		for _, m := range seedToSoilMap {
			if seed >= m.Start && seed <= (m.Start+m.Step) {
				soil = m.End + (seed - m.Start)
				fmt.Println(soil)
				break
			}
		}
		var fertilizer int
		for _, m := range soilToFertilizerMap {
			if soil >= m.Start && soil <= (m.Start+m.Step) {
				fertilizer = m.End + (soil - m.Start)
				fmt.Println(fertilizer)
				break
			}
		}
		var water int
		for _, m := range fertilizerToWaterMap {
			if fertilizer >= m.Start && fertilizer <= (m.Start+m.Step) {
				water = m.End + (fertilizer - m.Start)
				fmt.Println(water)
				break
			}
		}
	}
}
