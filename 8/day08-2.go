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
	lines := strings.Split(s, "\n")
	lines = lines[:len(lines)-1]

	var outputs []string
	var inputs []string
	total := 0
	for _, raw := range lines {
		matches := make([]string, 10)
		line := strings.Split(raw, "|")
		outputs = strings.Split(line[1], " ")[1:]

		inputs = strings.Split(line[0], " ")
		inputs = inputs[:len(inputs)-1]

		// Sort by length for ease
		sort.Slice(inputs, func(a, b int) bool { return len(inputs[a]) < len(inputs[b]) })

		// Get unique numbers
		matches[1] = inputs[0]
		matches[7] = inputs[1]
		matches[4] = inputs[2]
		matches[8] = inputs[9]

		// Get 6, 0, and 9
		for _, str_len_6 := range inputs[6:9] {
			contains_rune_1 := 0
			contains_rune_4 := 0
			for _, c := range matches[1] {
				if strings.ContainsRune(str_len_6, c) {
					contains_rune_1 += 1
				}
			}
			for _, c := range matches[4] {
				if strings.ContainsRune(str_len_6, c) {
					contains_rune_4 += 1
				}
			}
			if contains_rune_1 == 1 {
				matches[6] = str_len_6
			}
			if contains_rune_4 == 4 {
				matches[9] = str_len_6
			}
		}
		for _, str_len_6 := range inputs[6:9] {
			if strings.Compare(str_len_6, matches[6]) != 0 && strings.Compare(str_len_6, matches[9]) != 0 {
				matches[0] = str_len_6
			}
		}

		// Get 0, 2, 5
		for _, str_len_5 := range inputs[3:6] {
			contains_rune_1 := 0
			contains_rune_6 := 0
			for _, c := range matches[1] {
				if strings.ContainsRune(str_len_5, c) {
					contains_rune_1 += 1
				}
			}

			for _, c := range matches[6] {
				if strings.ContainsRune(str_len_5, c) {
					contains_rune_6 += 1
				}
			}

			if contains_rune_1 == 2 {
				matches[3] = str_len_5
			} else if contains_rune_6 == 5 {
				matches[5] = str_len_5
			} else {
				matches[2] = str_len_5
			}
		}

		var num_s string = ""

		for _, digit := range outputs {
			for i, poss := range matches {
				matched_runes := 0
				if len(poss) == len(digit) {
					for _, c := range digit {
						if strings.ContainsRune(poss, c) {
							matched_runes += 1
						}
					}
					if matched_runes == len(digit) {
						num_s += strconv.Itoa(i)
					}
				}
			}
		}
		val, _ := strconv.Atoi(num_s)
		total += val
	}
	return total
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
