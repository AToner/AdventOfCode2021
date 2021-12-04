package day04

import (
	"andytoner.com/aoc2021/pkg/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
--- Day 4: Giant Squid ---

You're already almost 1.5km (almost a mile) below the surface of the ocean, already so deep that you can't see any
sunlight. What you can see, however, is a giant squid that has attached itself to the outside of your submarine.

Maybe it wants to play bingo?

Bingo is played on a set of boards each consisting of a 5x5 grid of numbers. Numbers are chosen at random, and the
chosen number is marked on all boards on which it appears. (Numbers may not appear on all boards.) If all numbers in
any row or any column of a board are marked, that board wins. (Diagonals don't count.)

The submarine has a bingo subsystem to help passengers (currently, you and the giant squid) pass the time. It
automatically generates a random order in which to draw numbers and a random set of boards (your puzzle input). For
example:

7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7

After the first five numbers are drawn (7, 4, 9, 5, and 11), there are no winners, but the boards are marked as
follows (shown here adjacent to each other to save space):

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7

After the next six numbers are drawn (17, 23, 2, 0, 14, and 21), there are still no winners:

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7

Finally, 24 is drawn:

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7

At this point, the third board wins because it has at least one complete row or column of marked numbers (in this case,
the entire top row is marked: 14 21 17 24 4).

The score of the winning board can now be calculated. Start by finding the sum of all unmarked numbers on that board;
in this case, the sum is 188. Then, multiply that sum by the number that was just called when the board won, 24, to
get the final score, 188 * 24 = 4512.

To guarantee victory against the giant squid, figure out which board will win first. What will your final score be if
you choose that board?
31424
*/

type Square struct {
	value  int
	marked bool
}

type Board struct {
	board [5][5]Square
}

func Part1(fileName string) int {
	var boards []*Board
	input := utils.ReadLines(fileName)

	drawNumbers := createNumbers(input[0], ",")

	for i := 2; i < len(input); i += 6 {
		board := createBoard(input[i : i+5])
		boards = append(boards, board)
	}

	for _, number := range drawNumbers {
		markBoards(boards, number)
		winners := checkForWin(boards)
		if len(winners) > 0 {
			fmt.Println("Winners:")
			for _, winner := range winners {
				return sumBoard(winner, number)
			}
		}
	}
	return 0
}

func createNumbers(input string, separator string) []int {
	var result []int
	trimmedString := strings.Join(strings.Fields(input), " ")
	for _, stringNumber := range strings.Split(trimmedString, separator) {
		number, err := strconv.Atoi(stringNumber)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		result = append(result, number)
	}
	return result
}

func createBoard(input []string) *Board {
	board := &Board{
		board: [5][5]Square{},
	}
	for rowNumber, row := range input {
		columnNumbers := createNumbers(row, " ")
		for columnNumber, value := range columnNumbers {
			board.board[rowNumber][columnNumber] = Square{
				value:  value,
				marked: false,
			}
		}
	}
	return board
}

func markBoards(boards []*Board, value int) {
	markBoard := func(board *Board) {
		for rowNumber, _ := range board.board {
			for columnNumber, _ := range board.board[rowNumber] {
				square := &board.board[rowNumber][columnNumber]
				if square.value == value {
					square.marked = true
				}
			}
		}
	}

	for _, board := range boards {
		markBoard(board)
	}
}

func checkForWin(boards []*Board) []*Board {
	var result []*Board

	checkRows := func(board *Board) bool {
		for rowIndex, _ := range board.board {
			complete := true
			for _, square := range board.board[rowIndex] {
				if !square.marked {
					complete = false
				}
			}
			if complete {
				return true
			}
		}
		return false
	}

	checkColumns := func(board *Board) bool {
		for columnIndex := 0; columnIndex < len(board.board[0]); columnIndex++ {
			complete := true
			for rowIndex := 0; rowIndex < len(board.board); rowIndex++ {
				square := board.board[rowIndex][columnIndex]
				if !square.marked {
					complete = false
				}
			}
			if complete {
				return true
			}
		}
		return false
	}

	for _, board := range boards {
		if checkRows(board) || checkColumns(board) {
			result = append(result, board)
		}
	}
	return result
}

func sumBoard(board *Board, number int) int {
	var result int
	for rowNumber, _ := range board.board {
		for columnNumber, _ := range board.board[rowNumber] {
			square := board.board[rowNumber][columnNumber]
			if !square.marked {
				result += square.value
			}
		}
	}
	return result * number
}
