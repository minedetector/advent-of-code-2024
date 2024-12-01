package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strings"
    "strconv"
)

func main() {
  file, _ := os.Open("input.txt")
  defer file.Close()

  scanner := bufio.NewScanner(file)
  var list1, list2 []int

  for scanner.Scan() {
      locations := strings.Fields(scanner.Text())
      x1, _ := strconv.Atoi(locations[0])
      x2, _ := strconv.Atoi(locations[1])
      list1 = append(list1, x1)
      list2 = append(list2, x2)
  }

  fmt.Println(test1(list1, list2))
  fmt.Println(test2(list1, list2))
}

func test1(list1 []int, list2 []int) int {
  sort.Ints(list1)
  sort.Ints(list2)

  sum := 0

  for i := 0; i < len(list1); i++ {
    sum += abs(list1[i] - list2[i])
  }

  return sum
}

func test2(list1 []int, list2 []int) int {
  occurrences := map[int]int{}

  for _, value := range list2 {
      occurrences[value]++
  }

  sum := 0

  for _, value := range list1 {
    sum += value * occurrences[value]
  }

  return sum
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
