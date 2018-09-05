package main

import (
	"fmt"
	"time"
)

var nums = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

func solveSudoku(board [][]byte)  {
	solve(board, 0, 0)
}

func solve(board [][]byte, i, j int) bool {
	if i == 9 {
		fmt.Println("====")
		return true
	}
	if j == 9 {
		fmt.Println("====")
		return solve(board, i + 1, 0)
	}

	if board[i][j] == '.' {
		for _, v := range nums {
			board[i][j] = v
			if isValid(board, i, j) {
				if solve(board, i, j+1) {
					return true
				}
				fmt.Println(board)
				time.Sleep(time.Second)
			}
		}
		board[i][j] = '.'
	} else {
		return solve(board, i, j+1)
	}
	fmt.Println("====")
	return false
}

func isValid(board [][]byte, i, j int) bool {
	for row := 0; row < 9; row ++ {
		if board[i][j] == board[row][j] {
			return false
		}
	}

	for col := 0; col < 9; col ++ {
		if board[i][j] == board[i][col] {
			return false
		}
	}

	x := i / 3 * 3
	y := j / 3 * 3

	for row := x; row < x + 3; row ++ {
		for col := y; col < y + 3; col ++ {
			if board[i][j] == board[row][col] {
				return false
			}
		}
	}

	return true
}


func modify(b []byte) {
	b[0] = 'z'
}

func main() {
	d := [][]byte{[]byte{'5','3','.','.','7','.','.','.','.'},{'6','.','.','1','9','5','.','.','.'},{'.','9','8','.','.','.','.','6','.'},{'8','.','.','.','6','.','.','.','3'},{'4','.','.','8','.','3','.','.','1'},{'7','.','.','.','2','.','.','.','6'},{'.','6','.','.','.','.','2','8','.'},{'.','.','.','4','1','9','.','.','5'},{'.','.','.','.','8','.','.','7','9'}}
	fmt.Println(d)
	solveSudoku(d)

	modify(nums)
	fmt.Println(nums)

}