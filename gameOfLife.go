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

func render(board [][]int) {

	// Loops through each row and creates a new slice of strings. Each
	// element is then checked to see if it is 'alive' or 'dead'
	for _, row := range board {
		renderRow := make([]string, len(row))
		for index, state := range row {
			if state == 1 {
				renderRow[index] = "#"
			} else {
				renderRow[index] = " "
			}
		}
		fmt.Println(renderRow)
	}

}
