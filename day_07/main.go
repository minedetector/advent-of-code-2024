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

	output := 0

	var evaluate func(target int, numbers []int, index int, currentResult int) bool
	evaluate = func(target int, numbers []int, index int, currentResult int) bool {
		if index >= len(numbers) {
			return currentResult == target
		}

		currentNumber := numbers[index]

		newResult := currentResult+currentNumber

		if newResult <= target {
			if evaluate(target, numbers, index+1, newResult) {
				return true
			}
		}

		newResult = currentResult*currentNumber
		if newResult <= target {
			if evaluate(target, numbers, index+1, newResult) {
				return evaluate(target, numbers, index+1, newResult)
			}
		}

		concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", currentResult, currentNumber))
		if concatenated <= target {
			return evaluate(target, numbers, index+1, concatenated)
		}

		return false
	}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ": ")
		target, _ := strconv.Atoi(line[0])
		numberStrings := strings.Fields(line[1])
		numbers := make([]int, len(numberStrings))

		for i, numberString := range numberStrings {
			numbers[i], _ = strconv.Atoi(numberString)
		}

		if evaluate(target, numbers, 0, 0) {
			output += target
		}
	}

	fmt.Println(output)
}

