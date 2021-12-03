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

	depth := 0
	horiz := 0

	for _, raw := range lines {
		if len(raw) > 0 {
			spl := strings.Split(raw, " ")
			dist, _ := strconv.Atoi(spl[1])

			if strings.HasPrefix(spl[0], "forward") {
				horiz += dist
			} else if strings.HasPrefix(spl[0], "up") {
				depth -= dist
			} else {
				depth += dist
			}
		}
	}
	return horiz * depth
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
