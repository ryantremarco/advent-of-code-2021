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

	count := countWindowedIncreases(input)
	fmt.Println(count)
}

func countWindowedIncreases(depths []int) (out int) {
	for i, depth := range depths[3:] {
		a := depths[i] + depths[i+1] + depths[i+2]
		b := depths[i+1] + depths[i+2] + depth
		if b > a {
			out++
		}
	}

	return
}
