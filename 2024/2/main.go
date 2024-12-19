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
    inc := false
    dec := false
    safe := true
    removed := false
    for i := 0; i < len(lineSlice); i++ {
        fmt.Print("i")
        fmt.Println(i)
        fmt.Print("line: ")
        fmt.Println(lineSlice)
        fmt.Print("currently safe: ")
        fmt.Println(safe)
        if i + 1 == len(lineSlice) {
            break
        }
        x := lineSlice[i]
        y := lineSlice[i+1]
        if x < y {
            inc = true
        }
        if x > y {
            dec = true
        }
        if inc && dec && !removed {
            removed = true
            lineSlice = append(lineSlice[:i], lineSlice[i+1:]...)
            i = 0
        } else if inc && dec && removed {
            safe = false
        }
        if y - x > 3 && !removed {
            removed = true
            lineSlice = append(lineSlice[:i], lineSlice[i+1:]...)
            i = 0
        } else if y - x > 3 && removed {
            safe = false
        }
        if x - y > 3 && !removed {
            removed = true
            lineSlice = append(lineSlice[:i], lineSlice[i+1:]...)
            i = 0
        } else if x - y > 3 && removed {
            safe = false
        }
        if x == y && !removed {
            removed = true
            lineSlice = append(lineSlice[:i], lineSlice[i+1:]...)
            i = 0
        } else if x == y && removed {
            safe = false
        }
        if !safe {
            break
        }
    }
    return safe
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
    count := calcReport("2024/2/example.txt")
    fmt.Println(count)
}

