package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func run(s string) interface{} {
	d := " 1  4 90 66 38"

	l := strings.Fields(d)
	fmt.Println(l)
	fmt.Println(len(l))
	return 0
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
