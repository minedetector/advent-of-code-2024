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

	var input []string

	for scanner.Scan() {
		input = strings.Split(scanner.Text(), "")
	}

	blocks := []rune{}
	fileId := 0

	for i, numString := range input {
		num, _ := strconv.Atoi(numString)

		if i % 2 == 0 {
			for j := 0; j < num; j++ {
				blocks = append(blocks, rune('0'+fileId))
			}
			fileId++
		} else {
			for j := 0; j < num; j++ {
				blocks = append(blocks, '.')
			}
		}
	}

	lastNumber := len(blocks) - 1

	for i, char := range blocks {
		if i > lastNumber {
			break
		}
		if string(char) == "." {
			blocks[i] = blocks[lastNumber]
			blocks[lastNumber] = '.'
			lastNumber--
			for blocks[lastNumber] == '.' {
				lastNumber--
			}
		}
	}

	output := 0
	blockId := 0
	for blocks[blockId] != '.' {
		output += blockId * int(blocks[blockId]-'0')
		blockId++
	}

	fmt.Println("Output Day 9 Part 1", output)
}
