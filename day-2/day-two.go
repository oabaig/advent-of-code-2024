package main

import (
	"advent-of-code-2024/helpers"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

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
		var prev int
		var shouldIncrease bool
		for j := range len(reports[i]) {
			curr := helpers.StringToInt(reports[i][j])
			if j > 0 {
				// on the first element, determine whether the reports or increasing or decreasing
				if j == 1 {
					shouldIncrease = curr > prev
				} else {
					// check to see if the code is increasing or not when it should/should not
					if (shouldIncrease && prev >= curr) || (!shouldIncrease && curr >= prev) {
						break
					}
				}

				diff := int(math.Abs(float64(curr - prev)))

				// if the diff is not at least one and more than 3, then break out of the loop
				if diff < 1 || diff > 3 {
					break
				}
			}

			// if the code made it here, then it's a safe report
			if j == len(reports[i])-1 {
				numSafeReports++
			}
			prev = curr
		}
	}

	fmt.Println(numSafeReports)
}
