package main

import (
	"aoc/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := util.ReadInputStrings(os.Args[1])
	if err != nil {
		fmt.Printf("Failed to read input file: %s\n", err)
		os.Exit(1)
	}

	powerConsumption, err := calculatePowerConsumption(input)
	if err != nil {
		fmt.Printf("Failed to calculate power consumption: %s\n", err)
	}

	lifeSupport, err := calculateLifeSupportRating(input)
	if err != nil {
		fmt.Printf("Failed to calculate life support: %s\n", err)
	}

	fmt.Printf("puzzle1: %d\npuzzle 2: %d\n", powerConsumption, lifeSupport)
}

type diagnostics []string

// Returns the most common value as an int, or -1 if they are equally common
func (d diagnostics) commonBitAtIndex(i int) int {
	summed := 0
	for _, diagnostic := range d {
		if diagnostic[i] == '1' {
			summed++
		}
	}

	if summed == len(d)/2 {
		return -1
	}

	if summed > len(d)/2 {
		return 1
	}

	return 0
}

// puzzle 1
func calculatePowerConsumption(d diagnostics) (int64, error) {
	gammaS, epsilonS := gammaAndEpsilon(d)

	gamma, err := strconv.ParseInt(gammaS, 2, 64)
	if err != nil {
		return 0, err
	}

	epsilon, err := strconv.ParseInt(epsilonS, 2, 64)
	if err != nil {
		return 0, err
	}

	return gamma * epsilon, nil
}

func gammaAndEpsilon(d diagnostics) (string, string) {
	bitPerDiagnostic := len(d[0])
	var gammaSB strings.Builder
	var epsilonSB strings.Builder

	for i := 0; i < bitPerDiagnostic; i++ {
		common := d.commonBitAtIndex(i)

		if common == 1 || common == -1 {
			gammaSB.WriteRune('1')
			epsilonSB.WriteRune('0')
		} else {
			gammaSB.WriteRune('0')
			epsilonSB.WriteRune('1')
		}
	}

	return gammaSB.String(), epsilonSB.String()
}

// puzzle 2
func calculateLifeSupportRating(d diagnostics) (int64, error) {
	scrubberRatings := d
	oxygenRatings := d
	for i := 0; i < len(d[0]); i++ {
		ones := diagnostics{}
		zeros := diagnostics{}
		for _, diag := range oxygenRatings {
			bit := ([]rune(diag))[i]
			if bit == '1' {
				ones = append(ones, diag)
			} else {
				zeros = append(zeros, diag)
			}
		}

		if len(ones) >= len(zeros) {
			oxygenRatings = ones
		} else {
			oxygenRatings = zeros
		}
	}

	for i := 0; i < len(d[0]); i++ {
		ones := diagnostics{}
		zeros := diagnostics{}
		if len(scrubberRatings) == 1 {
			break
		}
		for _, diag := range scrubberRatings {
			bit := ([]rune(diag))[i]
			if bit == '1' {
				ones = append(ones, diag)
			} else {
				zeros = append(zeros, diag)
			}
		}

		if len(ones) < len(zeros) {
			scrubberRatings = ones
		} else {
			scrubberRatings = zeros
		}
	}

	oxygenRatingDecimal, err := strconv.ParseInt(oxygenRatings[0], 2, 64)
	if err != nil {
		return 0, err
	}

	scrubberRatingDecimal, err := strconv.ParseInt(scrubberRatings[0], 2, 64)
	if err != nil {
		return 0, err
	}

	return oxygenRatingDecimal * scrubberRatingDecimal, nil
}
