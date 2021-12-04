package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func splitSets(l []string, i int) (string, []string, []string) {
	cnt := 0
	var posOn []string
	var posOff []string
	for _, raw := range l {
		if string(raw[i]) == "1" {
			cnt += 1
			posOn = append(posOn, raw)
		} else {
			posOff = append(posOff, raw)
		}
	}

	if float64(cnt) >= (float64(len(l)) / 2.0) {
		return "1", posOn, posOff
	}
	return "0", posOn, posOff
}

func reduceWorkingSet(l []string) int64 {
	var oxygenSet []string
	var scrubberSet []string
	mcb, tmpSet1, tmpSet2 := splitSets(l, 0)

	if mcb == "1" {
		oxygenSet = tmpSet1
		scrubberSet = tmpSet2
	} else {
		oxygenSet = tmpSet2
		scrubberSet = tmpSet1
	}

	ind := 1
	for len(oxygenSet) > 1 {
		mcb, tmpSet1, tmpSet2 = splitSets(oxygenSet, ind)
		if mcb == "1" {
			oxygenSet = tmpSet1
		} else {
			oxygenSet = tmpSet2
		}
		ind += 1
	}

	ind = 1
	for len(scrubberSet) > 1 {
		mcb, tmpSet1, tmpSet2 = splitSets(scrubberSet, ind)
		if mcb == "1" {
			scrubberSet = tmpSet2
		} else {
			scrubberSet = tmpSet1
		}
		ind += 1
	}

	oxygen, _ := strconv.ParseInt(oxygenSet[0], 2, 64)
	scrubber, _ := strconv.ParseInt(scrubberSet[0], 2, 64)
	return oxygen * scrubber
}

func run(s string) interface{} {
	lines := strings.Split(s, "\n")
	lines = lines[:len(lines)-1]

	return reduceWorkingSet(lines)
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
