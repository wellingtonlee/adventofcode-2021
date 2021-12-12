package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func run(s string) interface{} {
	lines := strings.Split(s, "\n")
	lines = lines[:len(lines)-1]
	risk_sum := 0

	heightmap := make([][]int, len(lines))
	for i := range heightmap {
		heightmap[i] = make([]int, len(lines[0]))
		for j := range lines[i] {
			val, _ := strconv.Atoi(string(lines[i][j]))
			heightmap[i][j] = val
		}
	}

	for i, row := range heightmap {
		for j := range row {
			lower_than := 0
			if j == 0 {
				lower_than += 1
			} else if row[j] < row[j-1] {
				lower_than += 1
			}

			if j == len(row)-1 {
				lower_than += 1
			} else if row[j] < row[j+1] {
				lower_than += 1
			}

			if i == 0 {
				lower_than += 1
			} else if row[j] < heightmap[i-1][j] {
				lower_than += 1
			}

			if i == len(heightmap)-1 {
				lower_than += 1
			} else if row[j] < heightmap[i+1][j] {
				lower_than += 1
			}
			if lower_than == 4 {
				risk_sum += 1 + row[j]
			}

		}
	}

	return risk_sum
}

func main() {

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	start := time.Now()
	result := run(string(input))

	fmt.Printf("Duration: %f\n", time.Now().Sub(start).Seconds()*1000)
	fmt.Print("Result: ")
	fmt.Println(result)
}
