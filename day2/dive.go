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

	h, d, err := dive(input)
	if err != nil {
		fmt.Printf("Unexpected error: %s", err)
		os.Exit(1)
	}

	hAim, dAim, err := diveWithAim(input)
	if err != nil {
		fmt.Printf("Unexpected error: %s", err)
		os.Exit(1)
	}

	multiplied := h * d
	multipliedAim := hAim * dAim

	fmt.Printf("puzzle1: %d\npuzzle 2: %d\n", multiplied, multipliedAim)
}

// puzzle 1
func dive(directions []string) (int, int, error) {
	horizontal := 0
	depth := 0

	for _, entry := range directions {
		split := strings.Split(entry, " ")

		dir := split[0]
		amount, err := strconv.Atoi(split[1])
		if err != nil {
			return 0, 0, err
		}

		switch dir {
		case "forward":
			horizontal += amount
		case "up":
			depth -= amount
		case "down":
			depth += amount
		}
	}

	return horizontal, depth, nil
}

// puzzle 2
func diveWithAim(directions []string) (int, int, error) {
	horizontal := 0
	depth := 0
	aim := 0

	for _, entry := range directions {
		split := strings.Split(entry, " ")

		dir := split[0]
		amount, err := strconv.Atoi(split[1])
		if err != nil {
			return 0, 0, err
		}

		switch dir {
		case "forward":
			horizontal += amount
			depth += aim * amount
		case "up":
			aim -= amount
		case "down":
			aim += amount
		}
	}

	return horizontal, depth, nil
}
