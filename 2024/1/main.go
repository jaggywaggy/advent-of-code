package main

import (
  "fmt"
  "os"
  //"io"
  "strconv"
  "bufio"
)

type LocationIds struct {
  X []int
  Y []int
}

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}

func generateLocationIds(file string) LocationIds {
  var tokensX []int
  var tokensY []int


  data, err := os.Open(file)
  checkErr(err)
  defer data.Close()
  scanner := bufio.NewScanner(data)
  scanner.Split(bufio.ScanWords)

  i := 0
  for scanner.Scan() {
    if i == 1 {
      i = 0
      num, err := strconv.Atoi(scanner.Text())
      checkErr(err)
      tokensY = append(tokensY, num)
    } else {
      i++
      num, err := strconv.Atoi(scanner.Text())
      checkErr(err)
      tokensX = append(tokensX, num)
    }
  }

  locationIds := LocationIds{tokensX, tokensY}
  return locationIds
}

func main() {
  file := "ids.txt"
  locationIds := generateLocationIds(file)
  fmt.Println(locationIds)
}

