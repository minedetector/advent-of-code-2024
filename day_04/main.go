package main

import (
	"bufio"
	"fmt"
	"os"
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

	xmasCounter := 0
	xmasCounter2 := 0

	for y, line := range grid {
		for x, _ := range line {
			if isXmas(grid, y, x, 0, 1)   { xmasCounter++ }
			if isXmas(grid, y, x, 0, -1)  { xmasCounter++ }
			if isXmas(grid, y, x, 1, 0)   { xmasCounter++ }
			if isXmas(grid, y, x, -1, 0)  { xmasCounter++ }
			if isXmas(grid, y, x, 1, 1)   { xmasCounter++ }
			if isXmas(grid, y, x, -1, 1)  { xmasCounter++ }
			if isXmas(grid, y, x, 1, -1)  { xmasCounter++ }
			if isXmas(grid, y, x, -1, -1) { xmasCounter++ }
			if isXMas2(grid, y, x) { xmasCounter2++ }
		}
	}

	fmt.Println("xmasCounter: ", xmasCounter)
	fmt.Println("xmasCounter2: ", xmasCounter2)
}

func isXmas(grid [][]string, startY int, startX int, dy int, dx int) bool {
	xmas := "XMAS"
	for depth := 0; depth < 4; depth++ {
		y := startY + depth * dy
		x := startX + depth * dx
		if y < 0 || y >= len(grid) { return false }
		if x < 0 || x >= len(grid) { return false}
		if grid[y][x] != string(xmas[depth]) { return false }
	}
	return true
}

func isXMas2(grid[][]string, y int, x int) bool {
	if y + 2 >= len(grid) { return false }
	if x + 2 >= len(grid[y]) { return false }
	if grid[y + 1][x + 1] != "A" { return false }

	c1 := grid[y][x]
	c2 := grid[y][x + 2]
	c3 := grid[y + 2][x + 2]
	c4 := grid[y + 2][x]

	return (c1 == "M" && c3 == "S" || c1 == "S" && c3 == "M") &&
		   (c2 == "M" && c4 == "S" || c2 == "S" && c4 == "M")
}
