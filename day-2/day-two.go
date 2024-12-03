package main

import (
	"advent-of-code-2024/helpers"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func checkReactorReport(report []string, numProblematic int) int {

	fmt.Println(report, numProblematic)

	if len(report) == 1 {
		return numProblematic
	}

	var prev int
	var shouldIncrease bool
	for i := range len(report) {
		curr := helpers.StringToInt(report[i])
		if i > 0 {
			// on the first element, determine whether the reports or increasing or decreasing
			if i == 1 {
				shouldIncrease = curr > prev
			} else {
				// check to see if the code is increasing or not when it should/should not
				if (shouldIncrease && prev >= curr) || (!shouldIncrease && curr >= prev) {
					newReport := append(report[:i], report[i+1:]...)
					return checkReactorReport(newReport, numProblematic+1)
				}
			}

			diff := int(math.Abs(float64(curr - prev)))
			// if the diff is not at least one and more than 3, then break out of the loop
			if diff < 1 || diff > 3 {
				newReport := append(report[:i], report[i+1:]...)
				return checkReactorReport(newReport, numProblematic+1)
			}
		}

		prev = curr
	}

	return numProblematic
}

func main() {

	inFile := os.Args[1]

	f, err := os.Open(inFile)
	helpers.HandleError(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var reports [][]string
	for scanner.Scan() {
		reports = append(reports, []string{})
		text := scanner.Text()

		reactorReports := strings.Fields(text)
		reports[len(reports)-1] = append(reports[len(reports)-1], reactorReports...)
	}

	numSafeReports := 0
	for i := range len(reports) {
		result := checkReactorReport(reports[i], 0)
		fmt.Println(" ")
		if result <= 1 {
			numSafeReports++
		}
	}

	fmt.Println(numSafeReports)
}
