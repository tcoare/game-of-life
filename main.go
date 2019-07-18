package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Generates random seed so that results are different on each run
	rand.Seed(time.Now().UnixNano())

	randomState := randomState(7, 6)
	render(randomState)

}

// Generates a board with zero valued elements of a specified width and height
func deadState(width, height int) [][]int {

	// Creates board which is a slice containing height number of slices
	board := make([][]int, height)

	// loop to append 0 in width elements in each height slice in board
	for i := 0; i < height; i++ {
		for x := 0; x < width; x++ {
			board[i] = append(board[i], 0)
		}
	}

	// board_state = [
	//	    [0,0,0,0,0,0,0],
	//	    [0,0,0,0,0,0,0],
	//	    [0,0,0,0,0,0,0],
	//	    [0,0,0,0,0,0,0],
	//	    [0,0,0,0,0,0,0],
	//	    [0,0,0,0,0,0,0]
	//			]

	return board
}

// Generates a random state for each cell in board which is passed
func randomState(width, height int) [][]int {

	// Uses previous function to create a board which is zero entried
	board := deadState(width, height)

	// Loops through board and replaces each 0 value with a random number
	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			board[x][y] = rand.Intn(2)
		}
	}

	// board_state = [
	//	    [0,0,1,0,1,0,0],
	//	    [0,1,1,1,0,0,0],
	//	    [0,0,0,0,0,1,0],
	//	    [1,0,0,0,1,0,0],
	//	    [0,1,0,1,0,0,0],
	//	    [0,1,0,0,1,0,1]
	//			]

	return board
}

// Prints the board state to terminal
func render(board [][]int) {

	// TODO fix top and bottom border to always be the length of the rows

	// Loops through board and prints a hash for alive cells otherwise prints white
	// space

	fmt.Println("+----------------+")
	for _, row := range board {
		fmt.Print("| ")
		for _, state := range row {
			if state == 1 {
				fmt.Print("# ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println(" |")
	}
	fmt.Println("+----------------+")
}
