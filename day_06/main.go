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

	positions := 0

	//y, x
	directions := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	var posX, posY, sizeX int

	sizeY := len(grid)

	for y, line := range grid {
		sizeX = len(line)
		for x, char := range line {
			if char == "^" {
				posY = y
				posX = x
			}
		}
	}

	direction := 0
	for posX >= 0 && posX < sizeX-1 && posY >= 0 && posY < sizeY-1 {
		grid[posY][posX] = "X"
		if grid[posY+1*directions[direction][0]][posX+1*directions[direction][1]] == "#" {
			if direction == 3 {
				direction = 0
			} else {
				direction++
			}
		}
		posX += 1 * directions[direction][1]
		posY += 1 * directions[direction][0]
	}

	fmt.Println(directions)
	fmt.Println(posX)
	fmt.Println(posY)
	fmt.Println(sizeX)
	fmt.Println(sizeY)
	for _, g := range grid {
		for _, char := range g {
			if char == "X" {
				positions++
			}
		}
		fmt.Println(g)
	}
	fmt.Println(positions)
}
