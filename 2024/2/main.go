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
    for i := 0; i < len(lineSlice) - 1; i++ {
        x := common.Abs(lineSlice[i])
        y := common.Abs(lineSlice[i+1])
        diff := common.Abs(x - y)

        // check diff
        if diff < 1 || diff > 3 {
            if isOriginalSlice {
                return split(lineSlice, i)
            } else {
                return false
            }
        }

        // check inc or dec increments
        if (inc && dec) {
            if isOriginalSlice {
                return split(lineSlice, i)
            } else {
                return false
            }
        }
    }
    return true
}

func split(lineSlice []int, startIndex int) bool {
    var endIndex int = startIndex + 1
    var newSlice []int
    if startIndex + 1 > len(lineSlice) - 1 {
        newSlice = append(lineSlice[:endIndex-1])
    } else {
        newSlice = append(lineSlice[:startIndex], lineSlice[endIndex:]...)
    }
    return isSafe(newSlice, false)
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

