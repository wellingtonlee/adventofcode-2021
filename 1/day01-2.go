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
	numIncrease := 0
	lines := strings.Split(s, "\n")

	for i := 0; i < len(lines)-4; i++ {
		// We only need to really compare two numbers
		// since each sliding window has overlapping nums
		first, _ := strconv.Atoi(lines[i])
		second, _ := strconv.Atoi(lines[i+3])

		if second > first {
			numIncrease += 1
		}
	}
	return numIncrease
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
