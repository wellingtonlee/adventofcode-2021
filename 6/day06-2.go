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
	times_str := strings.Split(strings.Split(s, "\n")[0], ",")

	timers_old := [7]int{0, 0, 0, 0, 0, 0, 0}
	timers_new := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := range times_str {
		t, _ := strconv.Atoi(times_str[i])
		timers_old[t] += 1
	}

	var loc_old, tmp_old int
	var loc_new, tmp_new int
	for day := 0; day < 256; day++ {
		loc_old = day % 7
		loc_new = day % 9

		tmp_old = timers_old[loc_old]
		tmp_new = timers_new[loc_new]

		timers_new[loc_new] = tmp_old + tmp_new
		timers_old[loc_old] += tmp_new
	}

	num_fish := 0

	for i := range timers_old {
		num_fish += timers_old[i]
	}

	for i := range timers_new {
		num_fish += timers_new[i]
	}

	return num_fish
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
