package main

import (
	"aoc/util"
	"fmt"
	"os"
)

func main() {
	input, err := util.ReadInputInts(os.Args[1])
	if err != nil {
		fmt.Printf("Failed to read input file: %s\n", err)
		os.Exit(1)
	}

	count := countIncreases(input)
	fmt.Println(count)
}

func countIncreases(depths []int) (out int) {
	for i, depth := range depths[1:] {
		if depth > depths[i] {
			out++
		}
	}

	return
}
