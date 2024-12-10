package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]string

	for scanner.Scan() {
		chars := strings.Split(scanner.Text(), "")
		grid = append(grid, chars)
	}

	//x, y
	directions := [][]int{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}

	var findPath func(row int, col int, height int, reachNine map[[2]int]bool)
	findPath = func(row int, col int, height int, reachNine map [[2]int]bool) {
		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) {
			return
		}

		num, _ := strconv.Atoi(grid[row][col])
		if num != height {
			return
		}

		currentPosition := [2]int{row, col}

		if height == 9 {
			reachNine[currentPosition] = true
			return
		}

		for _, dir := range directions {
			nextRow := row + dir[0]
			nextColumn := col + dir[1]
			findPath(nextRow, nextColumn, height+1, reachNine)
		}
	}

	output := 0

	for row, line := range grid {
		for col, char := range line {
			if char == "0" {
				reachNine := make(map[[2]int]bool)
				findPath(row, col, 0, reachNine)

				output += len(reachNine)
			}
		}
	}

	fmt.Println(output)
}
