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

	priority := make(map[int][]int)

	var rules [][]string
	var pageNumbers [][]string
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "|") {
			rules = append(rules, strings.Split(scanner.Text(), "|"))
		} else if strings.Contains(scanner.Text(), ",") {
			pageNumbers = append(pageNumbers, strings.Split(scanner.Text(), ","))
		}
	}

	for _, rule := range rules {
		key, _ := strconv.Atoi(rule[0])
		value, _ := strconv.Atoi(rule[1])
		priority[key] = append(priority[key], value)
	}

	var safePages [][]string
	var unsafePages [][]string

	for _, page := range pageNumbers {
		if checkPage(page, priority) {
			safePages = append(safePages, page)
		} else {
			unsafePages = append(unsafePages, page)
		}
	}

	safeSum := 0
	unsafeSum := 0

	for _, safePage := range safePages {
		num, _ := strconv.Atoi(safePage[len(safePage)/2])
		safeSum += num
	}

	for _, unsafePage := range unsafePages {
		rebornPage := fixUnsafePage(unsafePage, priority)
		unsafeNum, _ := strconv.Atoi(rebornPage[len(rebornPage)/2])
		unsafeSum += unsafeNum
	}


	fmt.Println("safePages: ", safePages)
	fmt.Println("unsafePages: ", unsafePages)
	fmt.Println("safeSum: ", safeSum)
	fmt.Println("unsafeSum: ", unsafeSum)
}

func checkPage(page []string, priority map[int][]int) bool {
	var previousNumbers []int
	for _, number := range page {
		pageNumber, _ := strconv.Atoi(number)
		if priorityMismatch(previousNumbers, pageNumber, priority) {
			return false
		}
		previousNumbers = append(previousNumbers, pageNumber)
	}
	return true
}

func fixUnsafePage(page []string, priority map[int][]int) []string {
	var previousNumbers []int

	for index, number := range page {
		pageNumber, _ := strconv.Atoi(number)
		if priorityMismatch(previousNumbers, pageNumber, priority) {
			safeSlice := make([]string, len(page))
			copy(safeSlice, page)

			page = fixUnsafePage(reorderSlice(safeSlice, index), priority)
		}
		previousNumbers = append(previousNumbers, pageNumber)
	}
	return page
}

func priorityMismatch(previousNumbers []int, pageNumber int, priority map[int][]int) bool {
	for _, previous := range previousNumbers {
		for _, priorityValue := range priority[pageNumber] {
			if previous == priorityValue {
				return true
			}
		}
	}
	return false
}

func reorderSlice(slice []string, index int) []string {
	result := make([]string, len(slice))
	copy(result, slice)

	result[index], result[index-1] = result[index-1], result[index]

	return result
}
