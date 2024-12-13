package utils

import (
	"bufio"
	"os"
)

func GetInputAsArrayOfStrings(filePath string) ([]string, error) {
	var output []string
	readFile, err := os.Open(filePath)
	if err != nil {
		return output, err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		output = append(output, fileScanner.Text())
	}
	readFile.Close()
	return output, nil
}

func GetInputAs2DArrayOfStrings(filePath string) ([][]string, error) {
	var output [][]string
	readFile, err := os.Open(filePath)
	if err != nil {
		return output, err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		var row []string
		line := fileScanner.Text()
		for _, char := range line {
			row = append(row, string(char))
		}
		output = append(output, row)
	}
	readFile.Close()
	return output, nil
}

func GetInputAs2DArrayOfInts(filePath string) ([][]int, error) {
	var output [][]int
	readFile, err := os.Open(filePath)
	if err != nil {
		return output, err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		var row []int
		line := fileScanner.Text()
		for _, char := range line {
			row = append(row, ConvertStringToInt(string(char)))
		}
		output = append(output, row)
	}
	readFile.Close()
	return output, nil
}
