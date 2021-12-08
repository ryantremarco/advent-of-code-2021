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

	school, err := NewLanternfishSchool(input[0])
	if err != nil {
		fmt.Printf("Failed to create school: %s\n", err)
		os.Exit(1)
	}

	for i := 0; i < 80; i++ {
		school = school.age()
	}
	eightyDays := len(school)

	for i := 0; i < 176; i++ {
		school = school.age()
	}

	fmt.Printf("puzzle1: %d\npuzzle 2: %d\n", eightyDays, len(school))
}

type Lanternfish int
type LanternfishSchool []Lanternfish

func NewLanternfishSchool(input string) (LanternfishSchool, error) {
	var school LanternfishSchool
	asString := strings.Split(input, ",")
	for _, fishString := range asString {
		fish, err := strconv.Atoi(fishString)
		if err != nil {
			return nil, err
		}
		school = append(school, Lanternfish(fish))
	}
	return school, nil
}

func (l LanternfishSchool) age() LanternfishSchool {
	var newSchool LanternfishSchool
	for _, fish := range l {
		if fish == 0 {
			newSchool = append(newSchool, Lanternfish(8))
			fish = 7
		}
		newSchool = append(newSchool, fish-1)
	}
	return newSchool
}
