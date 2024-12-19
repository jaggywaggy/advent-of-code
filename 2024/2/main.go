package main

import (
  "os"
  "advent-of-code/common"
  "bufio"
  "fmt"
  "strings"
  "strconv"
)

func parseLineToIntSlice(line string) []int {
    var intSlice []int
    fields := strings.Fields(line)
    for _, field := range fields {
        num, err := strconv.Atoi(field)
        intSlice = append(intSlice, num)
        common.CheckErr(err)
    }
    return intSlice
}

func isSafe(lineSlice []int) bool {
    for i := 0; i < len(lineSlice); i++ {
        if i + 1 == len(lineSlice) {
            break
        }
        x := lineSlice[i]
        y := lineSlice[i+1]
        if x < y && y - x > 3 {
            return false
        }
        if y < x && x - y > 3 {
            return false
        }
    }
    return true
}

func calcReport(file string) int {
  data, err := os.Open(file)
  common.CheckErr(err)
  defer data.Close()
  scanner := bufio.NewScanner(data)
  count := 0
  for scanner.Scan() {
      line := scanner.Text()
      lineSlice := parseLineToIntSlice(line)
      if isSafe(lineSlice) {
          count++
      }
  }
  common.CheckErr(scanner.Err())
  return count
}

func main() {
    count := calcReport("2024/2/reports.txt")
    fmt.Println(count)
}

