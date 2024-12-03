package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pattern := `mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`

	sum1 := 0
	sum2 := 0

	increaseMemory := true

	for scanner.Scan() {
		code := scanner.Text()

		re := regexp.MustCompile(pattern)
		matches := re.FindAllStringSubmatch(code, -1)

		for _, match := range matches {
			if match[0] == "do()" {
				increaseMemory = true
			} else if match[0] == "don't()" {
				increaseMemory = false
			} else {
				// mul() match
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])

				sum1 += num1 * num2
				if increaseMemory {
					sum2 += num1 * num2
				}
			}
		}
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}
