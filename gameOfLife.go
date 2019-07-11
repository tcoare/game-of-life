package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// deadState := deadState(7, 6)
	// render(deadState)

	randomState := randomState(7, 6)
	render(randomState)

}

func deadState(width, height int) [][]int {

	board := make([][]int, height)

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

	board := deadState(width, height)

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
