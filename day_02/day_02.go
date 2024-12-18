package main

import (
    "bufio"
    "fmt"
    "os"
    "reflect"
    "sort"
    "strconv"
    "strings"
)

func main() {
    file, _ := os.Open("input.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)

    safeReports1 := 0
    safeReports2 := 0

    // Solution 1
    for scanner.Scan() {
        numbers := strings.Fields(scanner.Text())
        code := make([]int, len(numbers))
        for i, number := range numbers {
            num, err := strconv.Atoi(number)
            if err != nil {
                fmt.Printf("Error converting %s to integer: %v\n", number, err)
                break
            }
            code[i] = num
        }
        if isSafe(code) {
            safeReports1++
            safeReports2++
        } else if isSafeWithDampener(code) {
            safeReports2++
        } else {
            fmt.Println(code)
        }
    }

    fmt.Println("safeReports1: ", safeReports1)
    fmt.Println("safeReports2: ", safeReports2)
}

func isSafe(code []int) bool {
    ascending := make([]int, len(code))
    copy(ascending, code)
    sort.Ints(ascending)

    descending := make([]int, len(code))
    copy(descending, code)
    sort.Sort(sort.Reverse(sort.IntSlice(descending)))
    if reflect.DeepEqual(code, ascending) || reflect.DeepEqual(code, descending) {
        for i := 0; i<len(code)-1; i++ {
            x := abs(code[i] - code[i+1])
            if x == 0 || x > 3 {
                return false
            }
        }
    } else {
        return false
    }

    return true
}

func isSafeWithDampener(code []int) bool {
    for i := 0; i<len(code); i++ {
        orderChecker := RemoveIndex(code, i)

        if isSafe(orderChecker) {
            //fmt.Println("Correct: ", code)
            return true
        }
    }
    return false
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func RemoveIndex(s []int, index int) []int {
    result := make([]int, 0, len(s)-1)
    result = append(result, s[:index]...)
    return append(result, s[index+1:]...)
}
