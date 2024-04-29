package main

import (
	"fmt"
	"os"
)

func isValid(s *[9][9]int, row, col, val int) bool {
	for i := 0; i < 9; i++ {
		if s[row][i] == val || s[i][col] == val {
			return false
		}
	}

	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if s[i][j] == val {
				return false
			}
		}
	}
	return true
}

func solve(s *[9][9]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if s[row][col] == 0 {
				for val := 1; val <= 9; val++ {
					if isValid(s, row, col, val) {
						s[row][col] = val
						if solve(s) {
							return true
						}
						s[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func print(s *[9][9]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d ", s[i][j])
		}
		fmt.Println()
	}
}

func main() {
	if len(os.Args) != 9+1 {
		fmt.Println("Error")
		return
	}

	var sudoku [9][9]int
	for i := 0; i < 9; i++ {
		for j, char := range os.Args[1+i] {
			if char == '.' {
				sudoku[i][j] = 0
			} else {
				digit := int(char - '0')
				if digit < 1 || digit > 9 {
					fmt.Println("Error")
					return
				}
				sudoku[i][j] = digit
			}
		}
	}

	if !solve(&sudoku) {
		fmt.Println("Error")
	} else {
		print(&sudoku)
	}
}
