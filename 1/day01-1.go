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
	prev, _ := strconv.Atoi(lines[0])
	depth := 0

	for _, raw := range lines[1:] {
		depth, _ = strconv.Atoi(raw)
		if depth > prev {
			numIncrease += 1
		}
		prev = depth
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
