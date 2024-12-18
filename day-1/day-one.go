package main

import (
	"advent-of-code-2024/helpers"
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type InputData struct {
	first  []string
	second []string
}

func main() {

	inFile := os.Args[1]

	f, err := os.Open(inFile)
	helpers.HandleError(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	inputData := InputData{}

	for scanner.Scan() {
		text := scanner.Text()

		words := strings.Fields(text)

		inputData.first = append(inputData.first, words[0])
		inputData.second = append(inputData.second, words[1])
	}

	slices.Sort(inputData.first)
	slices.Sort(inputData.second)

	diff := 0
	similarity := 0
	for i := range len(inputData.first) {
		num1 := helpers.StringToInt(inputData.first[i])
		num2 := helpers.StringToInt(inputData.second[i])

		diff += int(math.Abs(float64(num1 - num2)))

		duplicateCount := 0
		for j := range len(inputData.second) {
			if inputData.first[i] == inputData.second[j] {
				duplicateCount++
			}
		}

		similarity += num1 * duplicateCount
	}

	fmt.Println("difference:", diff)
	fmt.Println("similarity score:", similarity)
}
