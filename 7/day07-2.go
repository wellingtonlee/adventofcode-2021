package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func run(s string) interface{} {
	pos_str := strings.Split(strings.Split(s, "\n")[0], ",")
	positions := make([]int, len(pos_str))

	max_pos := 0
	for i := range pos_str {
		t, _ := strconv.Atoi(pos_str[i])
		positions[i] = t
		if max_pos < t {
			max_pos = t
		}
	}

	// Build a lookup table of all summorials for speed
	summorials := make([]int, max_pos+1)

	for i := 1; i < max_pos+1; i++ {
		summorials[i] = i + summorials[i-1]
	}

	min_fuel := -1
	total := 0
	for i := 0; i < max_pos; i++ {
		total = 0
		for _, num := range positions {
			total += summorials[int(math.Abs(float64(num-i)))]
		}
		if total < min_fuel || min_fuel == -1 {
			min_fuel = total
		}
	}

	return min_fuel
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
