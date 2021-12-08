package main

import (
	"aoc/util"
	"fmt"
	"os"
	"strings"
)

const (
	Top         = "T"
	TopLeft     = "TL"
	TopRight    = "TR"
	Middle      = "M"
	BottomLeft  = "BL"
	BottomRight = "BR"
	Bottom      = "B"
)

func main() {
	input, err := util.ReadInputStrings(os.Args[1])
	if err != nil {
		fmt.Printf("Failed to read input file: %s\n", err)
		os.Exit(1)
	}

	displays := ParseDisplays(input)

	fmt.Printf("puzzle1: %d\npuzzle 2: %d\n", displays.countOutputUniquePatterns(), displays.totalOutput())
}

func subtractString(from, sub string) string {
	var sb strings.Builder
from:
	for _, fr := range from {
		for _, sr := range sub {
			if fr == sr {
				continue from
			}
		}
		sb.WriteRune(fr)
	}
	return sb.String()
}

type SegmentsNumberMapping map[Digit]int

func (s SegmentsNumberMapping) Get(d Digit) int {
	for key, value := range s {
		if key.Equals(d) {
			return value
		}
	}
	return 0
}

type Digit string

func (d Digit) Sub(other Digit) Digit {
	return Digit(subtractString(string(d), string(other)))
}

func (d Digit) Equals(other Digit) bool {
	if len(d) != len(other) {
		return false
	}

	if d.Sub(other) == "" && other.Sub(d) == "" {
		return true
	}

	return false
}

type Digits []Digit

func ParseDigits(digitStrings string) (d Digits) {
	for _, s := range strings.Split(digitStrings, " ") {
		d = append(d, Digit(s))
	}
	return
}

func (d Digits) OfLen(n int) (out Digits) {
	for _, digit := range d {
		if len(digit) == n {
			out = append(out, digit)
		}
	}
	return
}

type Display struct {
	patterns Digits
	output   Digits
}

func (d Displays) totalOutput() (total int) {
	for _, display := range d {
		mapping := display.deduceSegmentsNumberMapping()
		total += mapping.Get(display.output[0]) * 1000
		total += mapping.Get(display.output[1]) * 100
		total += mapping.Get(display.output[2]) * 10
		total += mapping.Get(display.output[3]) * 1
	}
	return
}

func (d Display) deduceSegmentsNumberMapping() SegmentsNumberMapping {
	mapping := SegmentsNumberMapping{}

	var one Digit
	var two Digit
	var three Digit
	var four Digit
	var five Digit
	var six Digit
	var seven Digit
	var eight Digit
	var nine Digit
	var zero Digit

	for _, pattern := range d.patterns {
		switch len(pattern) {
		case 2:
			one = pattern
		case 4:
			four = pattern
		case 3:
			seven = pattern
		case 7:
			eight = pattern
		}
	}

	fourdiff := four.Sub(one)

	for _, pattern := range d.patterns.OfLen(5) {
		if len(pattern.Sub(one)) == len(pattern)-2 {
			three = pattern
			continue
		}
		if len(pattern.Sub(fourdiff)) == len(pattern)-2 {
			five = pattern
			continue
		}
		two = pattern
	}

	for _, pattern := range d.patterns.OfLen(6) {
		if len(pattern.Sub(four)) == len(pattern)-4 {
			nine = pattern
			continue
		}
		if len(pattern.Sub(fourdiff)) == len(pattern)-2 {
			six = pattern
			continue
		}
		zero = pattern
	}

	mapping[one] = 1
	mapping[two] = 2
	mapping[three] = 3
	mapping[four] = 4
	mapping[five] = 5
	mapping[six] = 6
	mapping[seven] = 7
	mapping[eight] = 8
	mapping[nine] = 9
	mapping[zero] = 0

	return mapping
}

type Displays []Display

func ParseDisplays(input []string) Displays {
	var displays Displays
	for _, line := range input {
		if line == "" {
			continue
		}
		split := strings.Split(line, " | ")
		displays = append(displays, Display{
			patterns: ParseDigits(split[0]),
			output:   ParseDigits(split[1]),
		})
	}
	return displays
}

func (d Displays) countOutputUniquePatterns() (count int) {
	for _, display := range d {
		for _, digit := range display.output {
			switch len(digit) {
			case 2, 4, 3, 7:
				count++
			}
		}
	}
	return
}
