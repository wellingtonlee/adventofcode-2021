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
	counts := make([]int, len(lines[0]), len(lines[0]))

	for _, raw := range lines {
		for ind, chr := range raw {
			if string(chr) == "1" {
				counts[ind] += 1
			}
		}
	}

	gammaRate := ""
	epsilonRate := ""

	for i := 0; i < len(counts); i++ {
		if counts[i] > (len(lines)-1)/2 {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}

	gr, _ := strconv.ParseInt(gammaRate, 2, 64)
	er, _ := strconv.ParseInt(epsilonRate, 2, 64)
	return gr * er
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
