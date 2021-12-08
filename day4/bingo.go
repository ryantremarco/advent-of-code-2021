package main

import (
	"aoc/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const boardGridSize = 5

func main() {
	input, err := util.ReadInputStrings(os.Args[1])
	if err != nil {
		fmt.Printf("Failed to read input file: %s\n", err)
		os.Exit(1)
	}

	bingoCalls := []int{}
	for _, callString := range strings.Split(input[0], ",") {
		call, err := strconv.Atoi(callString)
		if err != nil {
			fmt.Printf("Failed to read bingo calls: %s\n", err)
			os.Exit(1)
		}
		bingoCalls = append(bingoCalls, call)
	}

	boardStrings := []string{}
	for _, inputLine := range input[2:] {
		if inputLine == "" {
			continue
		}
		boardStrings = append(boardStrings, inputLine)
	}

	boards := []bingoBoard{}
	for lineNo := 0; lineNo < len(boardStrings); lineNo += boardGridSize {
		board, err := parseBoard(boardStrings[lineNo : lineNo+boardGridSize])
		if err != nil {
			fmt.Printf("Failed to parse board: %s\n", err)
		}
		boards = append(boards, board)
	}

	winner, winningCall := winningBoard(bingoCalls, boards)
	loser, lastWinningCall := losingBoard(bingoCalls, boards)

	fmt.Printf("puzzle1: %d\npuzzle 2: %d\n", winner.sumUnmarked()*winningCall, loser.sumUnmarked()*lastWinningCall)

}

type bingoBoard [][]bingoNumber

func parseBoard(rows []string) (bingoBoard, error) {
	board := bingoBoard{}
	for _, row := range rows {
		rowNumbers := []bingoNumber{}
		for _, entry := range strings.Split(row, " ") {
			if entry == "" {
				continue
			}
			number, err := strconv.Atoi(entry)
			if err != nil {
				return nil, err
			}
			rowNumbers = append(rowNumbers, bingoNumber{
				value: number,
			})
		}
		board = append(board, rowNumbers)
	}
	return board, nil
}

func (b *bingoBoard) mark(toMark int) {
	for i, col := range *b {
		for j, number := range col {
			if number.value == toMark {
				(*b)[i][j].isMarked = true
			}
		}
	}
}

func (b bingoBoard) hasWon() bool {
	rowScores := make([]int, boardGridSize)
	colScores := make([]int, boardGridSize)
	for i, col := range b {
		for j, number := range col {
			if number.isMarked {
				colScores[i]++
				rowScores[j]++
			}
		}
	}

	for _, score := range rowScores {
		if score == boardGridSize {
			return true
		}
	}

	for _, score := range colScores {
		if score == boardGridSize {
			return true
		}
	}

	return false
}

func (b bingoBoard) sumUnmarked() (total int) {
	for _, col := range b {
		for _, number := range col {
			if !number.isMarked {
				total += number.value
			}
		}
	}
	return
}

type bingoNumber struct {
	value    int
	isMarked bool
}

// puzzle 1
func winningBoard(bingoCalls []int, boards []bingoBoard) (bingoBoard, int) {
	for _, call := range bingoCalls {
		for _, board := range boards {
			board.mark(call)
			if board.hasWon() {
				return board, call
			}
		}
	}
	return nil, 0
}

//puzzle 2
func losingBoard(bingoCalls []int, boards []bingoBoard) (bingoBoard, int) {
	wins := 0
	for _, call := range bingoCalls {
		for _, board := range boards {
			if !board.hasWon() {
				board.mark(call)
				if board.hasWon() {
					wins++
					if wins == len(boards)-1 {
						return board, call
					}
				}
			}
		}
	}

	return nil, 0
}
