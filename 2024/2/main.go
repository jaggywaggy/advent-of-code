package main

import (
  "os"
  "advent-of-code/common"
  "bufio"
)

func loadFile(file string) {
  data, err := os.Open(file)
  common.CheckErr(err)
  defer dat.Close()
  scanner := bufio.NewScanner(data)
  scanner.Split(bufio.ScanWord)
}

func main() {
}
