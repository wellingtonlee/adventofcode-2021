package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func run(s string) interface{} {
	pos_str := strings.Split(strings.Split(s, "\n")[0], ",")

	positions := make([]int, len(pos_str))
	total := 0
	max_pos := 0
	for i := range pos_str {
		t, _ := strconv.Atoi(pos_str[i])
		positions[i] = t
		total += t
		if max_pos < t {
			max_pos = t
		}
	}

	sort.Slice(positions, func(a, b int) bool { return positions[a] < positions[b] })

	min_fuel := total
	marker := 1

	for i := 1; i < max_pos; i++ {
		for j := marker; j < len(positions); j++ {
			if positions[j] < i {
				marker += 1
			}
		}

		total -= (len(positions) - marker)
		total += (marker)

		if total < min_fuel {
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
