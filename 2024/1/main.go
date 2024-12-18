// pt 1.a
// given two lists of integers (locationIds)
// sort in ascending order
// calculate diff between each list item
// sum the resulting diff values
package main

import (
	"fmt"
	"os"
	//"io"
	"strconv"
	"bufio"
	"sort"
	"advent-of-code/common"
)

type LocationIds struct {
	X []int
	Y []int
}

func loadLocationIds(file string) LocationIds {
	var tokensX []int
	var tokensY []int

	// first, open our file.
	data, err := os.Open(file)
	common.CheckErr(err)
	// close file upon scope exit.
	defer data.Close()
	// next, create a scanner scanning for words (default).
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanWords)

	i := 0
	// scan and extract the scanned text.
	for scanner.Scan() {
		if i == 1 {
			// second number, append to Y.
			i = 0
			num, err := strconv.Atoi(scanner.Text()) 	// int to str.
			common.CheckErr(err)
			tokensY = append(tokensY, num)						// append slice.
		} else {
			// first number, append to Y.
			i++
			num, err := strconv.Atoi(scanner.Text()) 	// int to str.
			common.CheckErr(err)
			tokensX = append(tokensX, num)						// append slice.
		}
	}

	// return our newly appended slices struct.
	locationIds := LocationIds{tokensX, tokensY}
	return locationIds
}

func calculateDiff(x []int, y []int) []int {
	var result []int
	for i := 0; i < len(x); i++ {
		if x[i] < y[i] {
			result = append(result, y[i] - x[i])
			continue
		}
		result = append(result, x[i] - y[i])
	}
	return result
}

func main() {
	var diffSlice []int
	file := "ids.txt"
	locationIds := loadLocationIds(file)
	// now we have our two lists. We will need to sort them (ascending).
	sort.Slice(locationIds.X, func(i, j int) bool {
		return locationIds.X[i] < locationIds.X[j]
	})
	sort.Slice(locationIds.Y, func(i, j int) bool {
		return locationIds.Y[i] < locationIds.Y[j]
	})
	// gather our diffs
	diffSlice = calculateDiff(locationIds.X, locationIds.Y)
	// sum our diffs
	sum := common.SumOfIntSlice(diffSlice)
	fmt.Println(sum)
	// out: 2769675
}

