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

type Vent struct {
	x1, x2, y1, y2 int
	valid          bool
}

func run(s string) interface{} {
	lines := strings.Split(s, "\n")

	vents := make([]Vent, 1, len(lines)-1)
	max_x, max_y := 0, 0

	for _, raw := range lines[:len(lines)-1] {
		coords := strings.Split(raw, " -> ")
		a_vals := strings.Split(coords[0], ",")
		b_vals := strings.Split(coords[1], ",")

		x1, _ := strconv.Atoi(a_vals[0])
		x2, _ := strconv.Atoi(b_vals[0])
		y1, _ := strconv.Atoi(a_vals[1])
		y2, _ := strconv.Atoi(b_vals[1])

		if x1 == x2 || y1 == y2 {
			//fmt.Printf("x: (%d, %d) | y: (%d, %d)\n", x1, x2, y1, y2)
			vents = append(vents, Vent{x1, x2, y1, y2, true})
			if x1 > max_x {
				max_x = x1
			}
			if x2 > max_x {
				max_x = x2
			}
			if y1 > max_y {
				max_y = y1
			}
			if y2 > max_y {
				max_y = y2
			}
		}
	}
	max_x += 1
	max_y += 1

	var floor [][]int
	for i := 0; i < max_x; i++ {
		row := make([]int, max_y)
		floor = append(floor, row)
	}
	start, end := 0, 0

	for _, vent := range vents {
		//fmt.Printf("x: (%d, %d) | y: (%d, %d) | %t\n", vent.x1, vent.x2, vent.y1, vent.y2, vent.valid)
		if vent.valid {
			if vent.x1 == vent.x2 {
				start = int(math.Min(float64(vent.y1), float64(vent.y2)))
				end = int(math.Max(float64(vent.y1), float64(vent.y2)))

				for j := start; j <= end; j++ {
					floor[vent.x1][j] += 1
				}
			} else {
				start = int(math.Min(float64(vent.x1), float64(vent.x2)))
				end = int(math.Max(float64(vent.x1), float64(vent.x2)))

				for j := start; j <= end; j++ {
					floor[j][vent.y1] += 1
				}
			}
		}
	}

	num_overlap := 0

	for i := 0; i < max_x; i++ {
		for j := 0; j < max_y; j++ {
			if floor[i][j] > 1 {
				num_overlap += 1
			}
		}
	}

	return num_overlap
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
