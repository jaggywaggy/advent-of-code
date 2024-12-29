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

func isSafe(lineSlice []int, isOriginalSlice bool) bool {
  for i := 0; i < len(lineSlice); i++ {
    if i + 1 > len(lineSlice) {
      break
    }
    x := common.Abs(lineSlice[i])
    y := common.Abs(lineSlice[i+1])
    inc := false
    dec := false
    triggerSplit := false
    if x > y {
      inc = true
      if common.Abs(x - y) > 3 {
        triggerSplit = true
      }
    } else {
      dec = true
      if common.Abs(y - x) > 3 {
        triggerSplit = true
      }
    }
    if inc && dec {
      triggerSplit = true
    }
    if triggerSplit && isOriginalSlice {
      return split(lineSlice, x)
    } else {
      return false
    }
  }
  return true
}

func split(lineSlice []int, indexToRemove int) bool {
  lineSlice = append(lineSlice[:indexToRemove], lineSlice[indexToRemove+1:]...)
  return isSafe(lineSlice, false)
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
    if isSafe(lineSlice, true) {
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

