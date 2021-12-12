package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func run(s string) interface{} {
	lines := strings.Split(s, "\n")
	lines = lines[:len(lines)-1]

	cnt := 0

	for _, raw := range lines {
		outputs := strings.Split(strings.Split(raw, "|")[1], " ")

		for _, s := range outputs {
			if len(s) == 2 || len(s) == 3 || len(s) == 4 || len(s) == 7 {
				cnt += 1
			}
		}
	}
	return cnt
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
