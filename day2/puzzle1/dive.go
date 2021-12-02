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

	multiplied := h * d
	fmt.Println(multiplied)
}

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
