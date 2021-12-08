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

	pipes, err := parsePipes(input)
	if err != nil {
		fmt.Printf("Failed to parse pipes: %s\n", err)
		os.Exit(1)
	}

	pipeIsStraight := func(p Pipe) bool {
		return p.start.x == p.end.x || p.start.y == p.end.y
	}

	pipeIsDiagonal := func(p Pipe) bool {
		return math.Abs(float64(p.end.x-p.start.x)) == math.Abs(float64(p.end.y-p.start.y))
	}

	straightPipeMap := pipes.generateStraightPipeMap(pipeIsStraight)

	diagonalPipeMap := pipes.generateStraightPipeMap(func(p Pipe) bool {
		return pipeIsStraight(p) || pipeIsDiagonal(p)
	})

	fmt.Printf("puzzle1: %d\npuzzle 2: %d\n", straightPipeMap.countOverlaps(), diagonalPipeMap.countOverlaps())
}

type Point struct {
	x int
	y int
}

type Pipe struct {
	start Point
	end   Point
}

func parsePipe(input string) (Pipe, error) {
	var points []Point
	for _, point := range strings.Split(input, " -> ") {
		coords := strings.Split(point, ",")
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			return Pipe{}, err
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			return Pipe{}, err
		}
		points = append(points, Point{x: x, y: y})
	}

	if len(points) != 2 {
		return Pipe{}, fmt.Errorf("expected 2 points from input, but received %d", len(points))
	}

	return Pipe{start: points[0], end: points[1]}, nil
}

func (p Pipe) lerp() []Point {
	var points []Point
	xSteps := p.end.x - p.start.x
	xInc := 0
	if xSteps != 0 {
		xInc = xSteps / int(math.Abs(float64(xSteps)))
	}
	ySteps := p.end.y - p.start.y
	yInc := 0
	if ySteps != 0 {
		yInc = ySteps / int(math.Abs(float64(ySteps)))
	}

	incFunc := func(x, y int) (int, int) {
		return x + xInc, y + yInc
	}

	for x, y := p.start.x, p.start.y; x != p.end.x+xInc || y != p.end.y+yInc; x, y = incFunc(x, y) {
		points = append(points, Point{x: x, y: y})
	}

	return points
}

type Pipes []Pipe

func parsePipes(inputs []string) (Pipes, error) {
	var pipes Pipes
	for _, input := range inputs {
		if input == "" {
			continue
		}
		pipe, err := parsePipe(input)
		if err != nil {
			return nil, err
		}
		pipes = append(pipes, pipe)
	}
	return pipes, nil
}

type PipeFilter func(Pipe) bool

func (p Pipes) generateStraightPipeMap(filter PipeFilter) PipeMap {
	var xBound int
	var yBound int
	for _, pipe := range p {
		if pipe.start.x+1 > xBound {
			xBound = pipe.start.x + 1
		}
		if pipe.end.x+1 > xBound {
			xBound = pipe.end.x + 1
		}
		if pipe.start.y+1 > yBound {
			yBound = pipe.start.y + 1
		}
		if pipe.end.y+1 > yBound {
			yBound = pipe.end.y + 1
		}
	}

	pipeMap := make(PipeMap, xBound)
	for i := range pipeMap {
		pipeMap[i] = make([]int, yBound)
		for j := range pipeMap[i] {
			pipeMap[i][j] = 0
		}
	}

	for _, pipe := range p {
		if filter(pipe) {
			for _, point := range pipe.lerp() {
				pipeMap[point.x][point.y]++
			}
		}
	}

	return pipeMap
}

type PipeMap [][]int

func (p PipeMap) countOverlaps() (count int) {
	for _, col := range p {
		for _, segment := range col {
			if segment >= 2 {
				count++
			}
		}
	}
	return
}
