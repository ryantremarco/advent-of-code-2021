package main

import (
	"aoc/util"
	"fmt"
	"math"
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

	crabs, err := parseCrabs(input[0])
	if err != nil {
		panic(err)
	}

	fmt.Printf("puzzle1: %d\npuzzle 2: %d\n", crabs.cheapestFuelToAlign(), crabs.cheapestIncrementalFuelToAlign())
}

func incrementalTo(x int) int {
	if x > 1 {
		return x + incrementalTo(x-1)
	}
	return 1
}

type Crab int

func (c Crab) fuelTo(x int) int {
	return int(math.Abs(float64(x - c.pos())))
}

func (c Crab) incrementalFuelTo(x int) int {
	to := int(math.Abs(float64(x - c.pos())))
	return incrementalTo(to)
}

func (c Crab) pos() int {
	return int(c)
}

type Crabs []Crab

func (c Crabs) fuelToAlignAt(x int) int {
	fuelCost := 0
	for _, crab := range c {
		fuelCost += crab.fuelTo(x)
	}
	return fuelCost
}

func (c Crabs) incrementalFuelToAlignAt(x int) int {
	var fuelCost int
	for _, crab := range c {
		fuelCost += crab.incrementalFuelTo(x)
	}
	return fuelCost
}

func (c Crabs) cheapestFuelToAlign() int {
	lowestX, highestX := c.extremePositions()
	cheapest := c.fuelToAlignAt(lowestX)
	for x := lowestX + 1; x <= highestX; x++ {
		if cost := c.fuelToAlignAt(x); cost < cheapest {
			cheapest = cost
		}
	}
	return cheapest
}

func (c Crabs) cheapestIncrementalFuelToAlign() int {
	lowestX, highestX := c.extremePositions()
	cheapest := c.incrementalFuelToAlignAt(lowestX)
	for x := lowestX + 1; x <= highestX; x++ {
		if cost := c.incrementalFuelToAlignAt(x); cost < cheapest {
			cheapest = cost
		}
	}
	return cheapest
}

func (c Crabs) extremePositions() (int, int) {
	lowest := c[0].pos()
	highest := c[0].pos()
	for _, crab := range c {
		pos := crab.pos()
		if pos > highest {
			highest = pos
		}
		if pos < lowest {
			lowest = pos
		}
	}
	return lowest, highest
}

func parseCrabs(input string) (Crabs, error) {
	var crabs Crabs
	crabStrings := strings.Split(input, ",")
	for _, crabString := range crabStrings {
		crab, err := strconv.Atoi(crabString)
		if err != nil {
			return nil, err
		}
		crabs = append(crabs, Crab(crab))
	}
	return crabs, nil
}
